package entity

import "encoding/xml"

type NFeProcRepository interface {
	SaveNFe(nfeProc *NFeProc) error
	DBPing() error
}

// type NFeProc struct
type NFeProc struct {
	XMLName xml.Name `xml:"nfeProc"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Versao  string   `xml:"versao,attr"`
	NFe     struct {
		Text   string `xml:",chardata"`
		InfNFe struct {
			Text   string `xml:",chardata"`
			Versao string `xml:"versao,attr"`
			ID     string `xml:"Id,attr"`
			Ide    struct {
				Text        string `xml:",chardata"`
				CUF         string `xml:"cUF"`
				CNF         string `xml:"cNF"`
				NatOp       string `xml:"natOp"`
				IndPag      string `xml:"indPag"`
				Mod         string `xml:"mod"`
				Serie       string `xml:"serie"`
				NNF         string `xml:"nNF"`
				DhEmi       string `xml:"dhEmi"`
				TpNF        string `xml:"tpNF"`
				IdDest      string `xml:"idDest"`
				CMunFG      string `xml:"cMunFG"`
				TpImp       string `xml:"tpImp"`
				TpEmis      string `xml:"tpEmis"`
				CDV         string `xml:"cDV"`
				TpAmb       string `xml:"tpAmb"`
				FinNFe      string `xml:"finNFe"`
				IndFinal    string `xml:"indFinal"`
				IndPres     string `xml:"indPres"`
				IndIntermed string `xml:"indIntermed"`
				ProcEmi     string `xml:"procEmi"`
				VerProc     string `xml:"verProc"`
				NFref       struct {
					Text   string `xml:",chardata"`
					RefNFe string `xml:"refNFe"`
				} `xml:"NFref"`
			} `xml:"ide"`
			Emit struct {
				Text      string `xml:",chardata"`
				CNPJ      string `xml:"CNPJ"`
				XNome     string `xml:"xNome"`
				XFant     string `xml:"xFant"`
				EnderEmit struct {
					Text    string `xml:",chardata"`
					XLgr    string `xml:"xLgr"`
					Nro     string `xml:"nro"`
					XBairro string `xml:"xBairro"`
					CMun    string `xml:"cMun"`
					XMun    string `xml:"xMun"`
					UF      string `xml:"UF"`
					CEP     string `xml:"CEP"`
					XPais   string `xml:"xPais"`
					Fone    string `xml:"fone"`
				} `xml:"enderEmit"`
				IE   string `xml:"IE"`
				IEST string `xml:"IEST"`
				IM   string `xml:"IM"`
				CNAE string `xml:"CNAE"`
				CRT  string `xml:"CRT"`
			} `xml:"emit"`
			Dest struct {
				Text      string `xml:",chardata"`
				CNPJ      string `xml:"CNPJ"`
				XNome     string `xml:"xNome"`
				EnderDest struct {
					Text    string `xml:",chardata"`
					XLgr    string `xml:"xLgr"`
					Nro     string `xml:"nro"`
					XBairro string `xml:"xBairro"`
					CMun    string `xml:"cMun"`
					XMun    string `xml:"xMun"`
					UF      string `xml:"UF"`
					CEP     string `xml:"CEP"`
					XPais   string `xml:"xPais"`
					Fone    string `xml:"fone"`
				} `xml:"enderDest"`
				IndIEDest string `xml:"indIEDest"`
				IE        string `xml:"IE"`
			} `xml:"dest"`
			Det []struct {
				Text  string `xml:",chardata"`
				NItem string `xml:"nItem,attr"`
				Prod  struct {
					Text     string `xml:",chardata"`
					CProd    string `xml:"cProd"`
					CEAN     string `xml:"cEAN"`
					XProd    string `xml:"xProd"`
					NCM      string `xml:"NCM"`
					CEST     string `xml:"CEST"`
					CFOP     string `xml:"CFOP"`
					UCom     string `xml:"uCom"`
					QCom     string `xml:"qCom"`
					VUnCom   string `xml:"vUnCom"`
					VProd    string `xml:"vProd"`
					CEANTrib string `xml:"cEANTrib"`
					UTrib    string `xml:"uTrib"`
					QTrib    string `xml:"qTrib"`
					VUnTrib  string `xml:"vUnTrib"`
					IndTot   string `xml:"indTot"`
					XPed     string `xml:"xPed"`
				} `xml:"prod"`
				Imposto struct {
					Text string `xml:",chardata"`
					ICMS struct {
						Text   string `xml:",chardata"`
						ICMS00 struct {
							Text  string `xml:",chardata"`
							Orig  string `xml:"orig"`
							CST   string `xml:"CST"`
							ModBC string `xml:"modBC"`
							VBC   string `xml:"vBC"`
							PICMS string `xml:"pICMS"`
							VICMS string `xml:"vICMS"`
						} `xml:"ICMS00"`
						ICMS10 struct {
							Text    string `xml:",chardata"`
							Orig    string `xml:"orig"`
							CST     string `xml:"CST"`
							ModBC   string `xml:"modBC"`
							VBC     string `xml:"vBC"`
							PICMS   string `xml:"pICMS"`
							VICMS   string `xml:"vICMS"`
							ModBCST string `xml:"modBCST"`
							PMVAST  string `xml:"pMVAST"`
							VBCST   string `xml:"vBCST"`
							PICMSST string `xml:"pICMSST"`
							VICMSST string `xml:"vICMSST"`
						} `xml:"ICMS10"`
					} `xml:"ICMS"`
					IPI struct {
						Text    string `xml:",chardata"`
						CEnq    string `xml:"cEnq"`
						IPITrib struct {
							Text string `xml:",chardata"`
							CST  string `xml:"CST"`
							VBC  string `xml:"vBC"`
							PIPI string `xml:"pIPI"`
							VIPI string `xml:"vIPI"`
						} `xml:"IPITrib"`
					} `xml:"IPI"`
					PIS struct {
						Text    string `xml:",chardata"`
						PISAliq struct {
							Text string `xml:",chardata"`
							CST  string `xml:"CST"`
							VBC  string `xml:"vBC"`
							PPIS string `xml:"pPIS"`
							VPIS string `xml:"vPIS"`
						} `xml:"PISAliq"`
					} `xml:"PIS"`
					COFINS struct {
						Text       string `xml:",chardata"`
						COFINSAliq struct {
							Text    string `xml:",chardata"`
							CST     string `xml:"CST"`
							VBC     string `xml:"vBC"`
							PCOFINS string `xml:"pCOFINS"`
							VCOFINS string `xml:"vCOFINS"`
						} `xml:"COFINSAliq"`
					} `xml:"COFINS"`
				} `xml:"imposto"`
				InfAdProd string `xml:"infAdProd"`
			} `xml:"det"`
			Total struct {
				Text    string `xml:",chardata"`
				ICMSTot struct {
					Text       string `xml:",chardata"`
					VBC        string `xml:"vBC"`
					VICMS      string `xml:"vICMS"`
					VICMSDeson string `xml:"vICMSDeson"`
					VFCP       string `xml:"vFCP"`
					VBCST      string `xml:"vBCST"`
					VST        string `xml:"vST"`
					VFCPST     string `xml:"vFCPST"`
					VFCPSTRet  string `xml:"vFCPSTRet"`
					VProd      string `xml:"vProd"`
					VFrete     string `xml:"vFrete"`
					VSeg       string `xml:"vSeg"`
					VDesc      string `xml:"vDesc"`
					VII        string `xml:"vII"`
					VIPI       string `xml:"vIPI"`
					VIPIDevol  string `xml:"vIPIDevol"`
					VPIS       string `xml:"vPIS"`
					VCOFINS    string `xml:"vCOFINS"`
					VOutro     string `xml:"vOutro"`
					VNF        string `xml:"vNF"`
				} `xml:"ICMSTot"`
			} `xml:"total"`
			Transp struct {
				Text       string `xml:",chardata"`
				ModFrete   string `xml:"modFrete"`
				Transporta struct {
					Text   string `xml:",chardata"`
					CNPJ   string `xml:"CNPJ"`
					XNome  string `xml:"xNome"`
					IE     string `xml:"IE"`
					XEnder string `xml:"xEnder"`
					XMun   string `xml:"xMun"`
					UF     string `xml:"UF"`
				} `xml:"transporta"`
				Vol struct {
					Text  string `xml:",chardata"`
					Marca string `xml:"marca"`
					NVol  string `xml:"nVol"`
					QVol  string `xml:"qVol"`
					PesoL string `xml:"pesoL"`
					PesoB string `xml:"pesoB"`
				} `xml:"vol"`
			} `xml:"transp"`
			Cobr struct {
				Text string `xml:",chardata"`
				Fat  struct {
					Text  string `xml:",chardata"`
					NFat  string `xml:"nFat"`
					VOrig string `xml:"vOrig"`
					VDesc string `xml:"vDesc"`
					VLiq  string `xml:"vLiq"`
				} `xml:"fat"`
			} `xml:"cobr"`
			Pag struct {
				Text   string `xml:",chardata"`
				DetPag struct {
					Text   string `xml:",chardata"`
					IndPag string `xml:"indPag"`
					TPag   string `xml:"tPag"`
					VPag   string `xml:"vPag"`
				} `xml:"detPag"`
			} `xml:"pag"`
			InfAdic struct {
				Text   string `xml:",chardata"`
				InfCpl string `xml:"infCpl"`
			} `xml:"infAdic"`
		} `xml:"infNFe"`
		Signature struct {
			Text       string `xml:",chardata"`
			Xmlns      string `xml:"xmlns,attr"`
			SignedInfo struct {
				Text                   string `xml:",chardata"`
				CanonicalizationMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"CanonicalizationMethod"`
				SignatureMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"SignatureMethod"`
				Reference struct {
					Text       string `xml:",chardata"`
					URI        string `xml:"URI,attr"`
					Transforms struct {
						Text      string `xml:",chardata"`
						Transform []struct {
							Text      string `xml:",chardata"`
							Algorithm string `xml:"Algorithm,attr"`
						} `xml:"Transform"`
					} `xml:"Transforms"`
					DigestMethod struct {
						Text      string `xml:",chardata"`
						Algorithm string `xml:"Algorithm,attr"`
					} `xml:"DigestMethod"`
					DigestValue string `xml:"DigestValue"`
				} `xml:"Reference"`
			} `xml:"SignedInfo"`
			SignatureValue string `xml:"SignatureValue"`
			KeyInfo        struct {
				Text     string `xml:",chardata"`
				X509Data struct {
					Text            string `xml:",chardata"`
					X509Certificate string `xml:"X509Certificate"`
				} `xml:"X509Data"`
			} `xml:"KeyInfo"`
		} `xml:"Signature"`
	} `xml:"NFe"`
	ProtNFe struct {
		Text    string `xml:",chardata"`
		Versao  string `xml:"versao,attr"`
		InfProt struct {
			Text     string `xml:",chardata"`
			TpAmb    string `xml:"tpAmb"`
			VerAplic string `xml:"verAplic"`
			ChNFe    string `xml:"chNFe"`
			DhRecbto string `xml:"dhRecbto"`
			NProt    string `xml:"nProt"`
			DigVal   string `xml:"digVal"`
			CStat    string `xml:"cStat"`
			XMotivo  string `xml:"xMotivo"`
		} `xml:"infProt"`
	} `xml:"protNFe"`
}
