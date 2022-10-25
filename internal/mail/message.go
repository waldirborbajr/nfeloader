package mail

import (
	"io"
	"os"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
	"github.com/rs/zerolog/log"
	"github.com/waldirborbajr/nfeloader/internal/entity"
)

// CustomerImapClient Call NewImapClient
func CustomerImapClient(usr, pwd, srv string) (*client.Client, error) {
	// 【Modify】Account and password
	return NewImapClient(usr, pwd, srv)
}

// NewImapClient Create an IMAP client
func NewImapClient(username, password, server string) (*client.Client, error) {
	// [Character set] When dealing with character sets other than us-ascii and utf-8 (such as gbk, gb2313, etc.),
	//  Need to add this line of code.
	// 【refer to】 https://github.com/emersion/go-imap/wiki/Charset-handling
	imap.CharsetReader = charset.Reader

	log.Info().Msg("Connecting to server...")

	// Connect to mail server
	c, err := client.DialTLS(server, nil)
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}
	log.Info().Msg("Connected")

	// Login with account password
	if err := c.Login(username, password); err != nil {
		return nil, err
	}

	log.Info().Msg("Logged in")

	return c, nil
}

func pop(list *[]uint32) uint32 {
	length := len(*list)
	lastEle := (*list)[length-1]
	*list = (*list)[:length-1]
	return lastEle
}

func NewMessage(path string, config *entity.NFeConfig) error {
	// Connect to mail server
	c, err := CustomerImapClient(config.MailUsr, config.MailPwd, config.MailServer)
	if err != nil {
		return err
	}
	// // Don't forget to logout
	// defer c.Logout()
	//
	// if err != nil {
	// 	return err
	// }

	// Select inbox
	_, err = c.Select("INBOX", false)
	if err != nil {
		return err
		// log.Fatal(err)
	}

	// Search condition instance object
	criteria := imap.NewSearchCriteria()

	// ALL is the default condition
	// See RFC 3501 section 6.4.4 for a list of searching criteria.
	criteria.WithoutFlags = []string{imap.SeenFlag}
	ids, _ := c.Search(criteria)
	var s imap.BodySectionName

	for {
		if len(ids) == 0 {
			break
		}
		id := pop(&ids)

		seqset := new(imap.SeqSet)
		seqset.AddNum(id)
		chanMessage := make(chan *imap.Message, 1)
		go func() {
			// The first fetch, only grab the message header, message flag, message size and other information, fast execution speed
			if err = c.Fetch(seqset,
				[]imap.FetchItem{imap.FetchEnvelope, imap.FetchFlags, imap.FetchRFC822Size}, chanMessage); err != nil {
				// [Practical experience] The err information encountered here is: ENVELOPE doesn't contain 10 fields
				// The reason is that the format of the email sent by the other party is not standardized, and the analysis fails
				// 相关的issue: https://github.com/emersion/go-imap/issues/143
				// log.Error().Msgf(seqset, err)

				log.Info().Msg(err.Error())
			}
		}()

		message := <-chanMessage
		if message == nil {
			log.Info().Msg("Server didn't returned message")
			continue
		}

		chanMsg := make(chan *imap.Message, 1)
		go func() {
			// This is the second fetch, to get the MIME content of the mail
			if err = c.Fetch(seqset,
				[]imap.FetchItem{imap.FetchRFC822},
				chanMsg); err != nil {
				log.Error().Err(err).Msg("Error fetching message")
			}
		}()

		msg := <-chanMsg
		if msg == nil {
			log.Info().Msg("Server didn't returned message")
		}

		section := &s
		r := msg.GetBody(section)
		if r == nil {
			// return r
			log.Info().Msg("Server didn't returned message body")
		}

		// Create a new mail reader
		mr, err := mail.CreateReader(r)
		if err != nil {
			return err
		}

		// Process each message's part
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}

			switch h := p.Header.(type) {
			case *mail.AttachmentHeader:
				// This is an attachment
				filename, err := h.Filename()
				if err != nil {
					return err
				}
				if filename != "" {
					log.Info().Msg("Got attachment: ")
					b, _ := io.ReadAll(p.Body)
					file, _ := os.OpenFile(path+filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
					defer file.Close()
					_, err := file.Write(b)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	// Don't forget to logout
	if err := c.Logout(); err != nil {
		return err
	}

	return nil
}
