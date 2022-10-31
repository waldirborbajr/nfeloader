package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/waldirborbajr/nfeloader/internal/version"
)

var tlsCertFile = "./certs/certbundle.pem"
var tlsKeyFile = "./certs/server.key"
var isTLS bool
var listenAddr string

type FileInfo struct {
	os.FileInfo
}

type FileInfoList struct {
	fileInfoList []os.FileInfo
}

type Server struct {
	ListenAddr string
}

func init() {
	isTLS, _ = strconv.ParseBool(os.Getenv("TLS"))

	if !isTLS {
		listenAddr = os.Getenv("LISTEN_ADDR")
	} else {
		listenAddr = os.Getenv("LISTEN_ADDR_TLS")
	}
}

func (f FileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Name":    f.Name(),
		"Size":    f.Size(),
		"Mode":    f.Mode(),
		"ModTime": f.ModTime(),
		"IsDir":   f.IsDir(),
	})
}

func (f FileInfoList) MarshalJSON() ([]byte, error) {
	fileInfoList := make([]FileInfo, 0, len(f.fileInfoList))
	for _, val := range f.fileInfoList {
		fileInfoList = append(fileInfoList, FileInfo{val})
	}

	return json.Marshal(fileInfoList)
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
	}
}

func (s *Server) Start() {
	log.Printf("NFeLoader API - HTTP server is running on port %s", s.ListenAddr)

	mux := http.NewServeMux()
	setupHandlers(mux)

	if isTLS {
		log.Fatal(http.ListenAndServeTLS(s.ListenAddr, tlsCertFile, tlsKeyFile, mux))
	} else {

		log.Fatal(http.ListenAndServe(s.ListenAddr, mux))
	}
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/xmls", handleXMLs)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("NFe Loader Audit! " + version.AppVersion()))
}

func handleXMLs(w http.ResponseWriter, r *http.Request) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error on Getwd -> %s", err.Error())
	}

	dir, err := os.Open(pwd + "/xmls")
	if err != nil {
		fmt.Printf("Error on Open -> %s", err.Error())
	}

	entries, err := dir.Readdir(0)
	if err != nil {
		fmt.Printf("Error on ReadDir -> %s", err.Error())
	}

	output, err := json.Marshal(FileInfoList{entries})
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	// default port on empty
	if len(listenAddr) == 0 {
		listenAddr = ":9191"
	}

	s := NewServer(listenAddr)

	s.Start()
}
