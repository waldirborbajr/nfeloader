package repository

import (
	"database/sql"

	"github.com/rs/zerolog/log"
	"github.com/waldirborbajr/nfeloader/pkg/customlog"
	"github.com/waldirborbajr/nfeloader/pkg/entity"
)

type NFeProcRepositoryMysql struct {
	db *sql.DB
}

func NewNFeProcMysql(db *sql.DB) *NFeProcRepositoryMysql {
	return &NFeProcRepositoryMysql{db: db}
}

func (n *NFeProcRepositoryMysql) DBPing() error {
	if err := n.db.Ping(); err != nil {
		return err
	}
	return nil
}

func (n *NFeProcRepositoryMysql) SaveNFe(nfeProc *entity.NFeProc) error {
	nNF := nfeProc.NFe.InfNFe.Ide.NNF

	log.Info().Msg("Preparing to insert into NFe")

	stmt, err := n.db.Prepare(
		`
		INSERT INTO nfe
		(
			ide_cUF,
			ide_cNF,
			ide_natOp,
			ide_indPag,
			ide_mod,
			ide_serie,
			ide_nNF,
			ide_dhEmi,
			ide_tpNF,
			ide_idDest,
			ide_cMunFG,
			ide_tpImp,
			ide_tpEmis,
			ide_cDV,
			ide_tpAmb,
			ide_finNFe,
			ide_indFinal,
			ide_indPres,
			ide_procEmi,
			ide_verProc,
			nfref_refNFe,
			emit_CNPJ,
			emit_xNome,
			emit_xFant,
			emit_xLgr,
			emit_nro,
			emit_xBairro,
			emit_cMun,
			emit_xMun,
			emit_UF,
			emit_CEP,
			emit_xPais,
			emit_fone,
			emit_IE,
			emit_IEST,
			emit_IM,
			emit_CNAE,
			emit_CRT,
			dest_CNPJ,
			dest_xNome,
			dest_xLgr,
			dest_nro,
			dest_xBairro,
			dest_cMun,
			dest_xMun,
			dest_UF,
			dest_CEP,
			dest_xPais,
			dest_fone,
			dest_indIEDest,
			dest_IE,
			total_vBC,
			total_vICMS,
			total_vICMSDeson,
			total_vFCP,
			total_vBCST,
			total_vST,
			total_vFCPST,
			total_vFCPSTRet,
			total_vProd,
			total_vFrete,
			total_vSeg,
			total_vDesc,
			total_vII,
			total_vIPI,
			total_vIPIDevol,
			total_vPIS,
			total_vCOFINS,
			total_vOutro,
			total_vNF,
			transp_modFrete,
			transp_CNPJ,
			transp_xNome,
			transp_IE,
			transp_xEnder,
			transp_xMun,
			transp_UF,
			transp_qVol,
			transp_marca,
			transp_nVol,
			transp_pesoL,
			transp_pesoB,
			cobr_nFat,
			cobr_vOrig,
			cobr_vDesc,
			cobr_vLiq,
			detPag_indPag,
			detPag_tPag,
			detPag_vPag,
			infAdic_infCpl
		)
		VALUES
		(
			?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,
			?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,
			?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?
		)
		ON DUPLICATE KEY UPDATE
		ide_cUF=VALUES(ide_cUF),
		ide_cNF=VALUES(ide_cNF),
		ide_natOp=VALUES(ide_natOp),
		ide_mod=VALUES(ide_mod),
		ide_indPag=VALUES(ide_indPag),
		ide_serie=VALUES(ide_serie),
		ide_nNF=VALUES(ide_nNF),
		ide_dhEmi=VALUES(ide_dhEmi),
		ide_tpNF=VALUES(ide_tpNF),
		ide_idDest=VALUES(ide_idDest),
		ide_cMunFG=VALUES(ide_cMunFG),
		ide_tpImp=VALUES(ide_tpImp),
		ide_tpEmis=VALUES(ide_tpEmis),
		ide_cDV=VALUES(ide_cDV),
		ide_tpAmb=VALUES(ide_tpAmb),
		ide_finNFe=VALUES(ide_finNFe),
		ide_indFinal=VALUES(ide_indFinal),
		ide_indPres=VALUES(ide_indPres),
		ide_procEmi=VALUES(ide_procEmi),
		ide_verProc=VALUES(ide_verProc),
		nfref_refNFe=VALUES(nfref_refNFe),
		emit_CNPJ=VALUES(emit_CNPJ),
		emit_xNome=VALUES(emit_xNome),
		emit_xFant=VALUES(emit_xFant),
		emit_xLgr=VALUES(emit_xLgr),
		emit_nro=VALUES(emit_nro),
		emit_xBairro=VALUES(emit_xBairro),
		emit_cMun=VALUES(emit_cMun),
		emit_xMun=VALUES(emit_xMun),
		emit_UF=VALUES(emit_UF),
		emit_CEP=VALUES(emit_CEP),
		emit_xPais=VALUES(emit_xPais),
		emit_fone=VALUES(emit_fone),
		emit_IE=VALUES(emit_IE),
		emit_IEST=VALUES(emit_IEST),
		emit_IM=VALUES(emit_IM),
		emit_CNAE=VALUES(emit_CNAE),
		emit_CRT=VALUES(emit_CRT),
		dest_CNPJ=VALUES(dest_CNPJ),
		dest_xNome=VALUES(dest_xNome),
		dest_xLgr=VALUES(dest_xLgr),
		dest_nro=VALUES(dest_nro),
		dest_xBairro=VALUES(dest_xBairro),
		dest_cMun=VALUES(dest_cMun),
		dest_xMun=VALUES(dest_xMun),
		dest_UF=VALUES(dest_UF),
		dest_CEP=VALUES(dest_CEP),
		dest_xPais=VALUES(dest_xPais),
		dest_fone=VALUES(dest_fone),
		dest_indIEDest=VALUES(dest_indIEDest),
		dest_IE=VALUES(dest_IE),
		total_vBC=VALUES(total_vBC),
		total_vICMS=VALUES(total_vICMS),
		total_vICMSDeson=VALUES(total_vICMSDeson),
		total_vFCP=VALUES(total_vFCP),
		total_vBCST=VALUES(total_vBCST),
		total_vST=VALUES(total_vST),
		total_vFCPST=VALUES(total_vFCPST),
		total_vFCPSTRet=VALUES(total_vFCPSTRet),
		total_vProd=VALUES(total_vProd),
		total_vFrete=VALUES(total_vFrete),
		total_vSeg=VALUES(total_vSeg),
		total_vDesc=VALUES(total_vDesc),
		total_vII=VALUES(total_vII),
		total_vIPI=VALUES(total_vIPI),
		total_vIPIDevol=VALUES(total_vIPIDevol),
		total_vPIS=VALUES(total_vPIS),
		total_vCOFINS=VALUES(total_vCOFINS),
		total_vOutro=VALUES(total_vOutro),
		total_vNF=VALUES(total_vNF),
		transp_modFrete=VALUES(transp_modFrete),
		transp_CNPJ=VALUES(transp_CNPJ),
		transp_xNome=VALUES(transp_xNome),
		transp_IE=VALUES(transp_IE),
		transp_xEnder=VALUES(transp_xEnder),
		transp_xMun=VALUES(transp_xMun),
		transp_UF=VALUES(transp_UF),
		transp_qVol=VALUES(transp_qVol),
		transp_marca=VALUES(transp_marca),
		transp_nVol=VALUES(transp_nVol),
		transp_pesoL=VALUES(transp_pesoL),
		transp_pesoB=VALUES(transp_pesoB),
		cobr_nFat=VALUES(cobr_nFat),
		cobr_vOrig=VALUES(cobr_vOrig),
		cobr_vDesc=VALUES(cobr_vDesc),
		cobr_vLiq=VALUES(cobr_vLiq),
		detPag_indPag=VALUES(detPag_indPag),
		detPag_tPag=VALUES(detPag_tPag),
		detPag_vPag=VALUES(detPag_vPag),
		infAdic_infCpl=VALUES(infAdic_infCpl)
		`)
	if err != nil {
		customlog.HandleError("Preparing NFe", err)
		return err
	}

	_, err = stmt.Exec(
		nfeProc.NFe.InfNFe.Ide.CUF,
		nfeProc.NFe.InfNFe.Ide.CNF,
		nfeProc.NFe.InfNFe.Ide.NatOp,
		nfeProc.NFe.InfNFe.Ide.IndPag,
		nfeProc.NFe.InfNFe.Ide.Mod,
		nfeProc.NFe.InfNFe.Ide.Serie,
		nfeProc.NFe.InfNFe.Ide.NNF,
		nfeProc.NFe.InfNFe.Ide.DhEmi,
		nfeProc.NFe.InfNFe.Ide.TpNF,
		nfeProc.NFe.InfNFe.Ide.IdDest,
		nfeProc.NFe.InfNFe.Ide.CMunFG,
		nfeProc.NFe.InfNFe.Ide.TpImp,
		nfeProc.NFe.InfNFe.Ide.TpEmis,
		nfeProc.NFe.InfNFe.Ide.CDV,
		nfeProc.NFe.InfNFe.Ide.TpAmb,
		nfeProc.NFe.InfNFe.Ide.FinNFe,
		nfeProc.NFe.InfNFe.Ide.IndFinal,
		nfeProc.NFe.InfNFe.Ide.IndPres,
		nfeProc.NFe.InfNFe.Ide.ProcEmi,
		nfeProc.NFe.InfNFe.Ide.VerProc,
		nfeProc.NFe.InfNFe.Ide.NFref.RefNFe,
		nfeProc.NFe.InfNFe.Emit.CNPJ,
		nfeProc.NFe.InfNFe.Emit.XNome,
		nfeProc.NFe.InfNFe.Emit.XFant,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.XLgr,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.Nro,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.XBairro,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.CMun,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.XMun,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.UF,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.CEP,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.XPais,
		nfeProc.NFe.InfNFe.Emit.EnderEmit.Fone,
		nfeProc.NFe.InfNFe.Emit.IE,
		nfeProc.NFe.InfNFe.Emit.IEST,
		nfeProc.NFe.InfNFe.Emit.IM,
		nfeProc.NFe.InfNFe.Emit.CNAE,
		nfeProc.NFe.InfNFe.Emit.CRT,
		nfeProc.NFe.InfNFe.Dest.CNPJ,
		nfeProc.NFe.InfNFe.Dest.XNome,
		nfeProc.NFe.InfNFe.Dest.EnderDest.XLgr,
		nfeProc.NFe.InfNFe.Dest.EnderDest.Nro,
		nfeProc.NFe.InfNFe.Dest.EnderDest.XBairro,
		nfeProc.NFe.InfNFe.Dest.EnderDest.CMun,
		nfeProc.NFe.InfNFe.Dest.EnderDest.XMun,
		nfeProc.NFe.InfNFe.Dest.EnderDest.UF,
		nfeProc.NFe.InfNFe.Dest.EnderDest.CEP,
		nfeProc.NFe.InfNFe.Dest.EnderDest.XPais,
		nfeProc.NFe.InfNFe.Dest.EnderDest.Fone,
		nfeProc.NFe.InfNFe.Dest.IndIEDest,
		nfeProc.NFe.InfNFe.Dest.IE,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VBC,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VICMS,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VICMSDeson,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VFCP,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VBCST,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VST,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VFCPST,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VFCPSTRet,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VProd,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VFrete,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VSeg,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VDesc,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VII,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VIPI,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VIPIDevol,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VPIS,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VCOFINS,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VOutro,
		nfeProc.NFe.InfNFe.Total.ICMSTot.VNF,
		nfeProc.NFe.InfNFe.Transp.ModFrete,
		nfeProc.NFe.InfNFe.Transp.Transporta.CNPJ,
		nfeProc.NFe.InfNFe.Transp.Transporta.XNome,
		nfeProc.NFe.InfNFe.Transp.Transporta.IE,
		nfeProc.NFe.InfNFe.Transp.Transporta.XEnder,
		nfeProc.NFe.InfNFe.Transp.Transporta.XMun,
		nfeProc.NFe.InfNFe.Transp.Transporta.UF,
		nfeProc.NFe.InfNFe.Transp.Vol.Marca,
		nfeProc.NFe.InfNFe.Transp.Vol.NVol,
		nfeProc.NFe.InfNFe.Transp.Vol.QVol,
		nfeProc.NFe.InfNFe.Transp.Vol.PesoL,
		nfeProc.NFe.InfNFe.Transp.Vol.PesoB,
		nfeProc.NFe.InfNFe.Cobr.Fat.NFat,
		nfeProc.NFe.InfNFe.Cobr.Fat.VOrig,
		nfeProc.NFe.InfNFe.Cobr.Fat.VDesc,
		nfeProc.NFe.InfNFe.Cobr.Fat.VLiq,
		nfeProc.NFe.InfNFe.Pag.DetPag.IndPag,
		nfeProc.NFe.InfNFe.Pag.DetPag.TPag,
		nfeProc.NFe.InfNFe.Pag.DetPag.VPag,
		nfeProc.NFe.InfNFe.InfAdic.InfCpl)

	if err != nil {
		customlog.HandleError("Exec of NFe", err)
		return err
	}

	log.Printf("Validating for previous NFeDetalhe")

	var exists bool
	row := n.db.QueryRow("SELECT EXISTS(SELECT 1 FROM nfedetalhe WHERE nNF = ?)", nNF)

	if err = row.Scan(&exists); err != nil {
		customlog.HandleError("Exist Detail ", err)
		return err
	}

	if exists {
		log.Info().Msg("Deleting NFeDetalhe")

		stmt, err = n.db.Prepare(`DELETE FROM nfedetalhe WHERE nNF = ?`)

		if err != nil {
			customlog.HandleError("Prepare of delete NFeDetalhe", err)
			return err
		}

		if _, err = stmt.Exec(nNF); err != nil {
			customlog.HandleError("Excuting delete NFeDetalhe", err)
			return err
		}
	}

	for _, detail := range nfeProc.NFe.InfNFe.Det {
		stmt, err = n.db.Prepare(
			`
			INSERT INTO nfedetalhe
			(
				nNF,
				det_cProd,
				det_cEAN,
				det_xProd,
				det_NCM,
				det_CEST,
				det_CFOP,
				det_uCom,
				det_qCom,
				det_vUnCom,
				det_vProd,
				det_cEANTrib,
				det_uTrib,
				det_qTrib,
				det_vUnTrib,
				det_indTot,
				det_xPed,
				det_ICMS_orig,
				det_ICMS_CST,
				det_ICMS_modBC,
				det_ICMS_vBC,
				det_ICMS_pICMS,
				det_ICMS_vICMS,
				det_ICMS_modBCST,
				det_ICMS_pMVAST,
				det_ICMS_vBCST,
				det_ICMS_pICMSST,
				det_ICMS_vICMSST,
				det_IPI_CST,
				det_IPI_vBC,
				det_IPI_pIPI,
				det_IPI_vIPI,
				det_PIS_CST,
				det_PIS_VBC,
				det_PIS_PPIS,
				det_PIS_VPIS,
				det_COFINS_CST,
				det_COFINS_VBC,
				det_COFINS_PCOFINS,
				det_COFINS_VCOFINS,
				det_infAdProd
			)
			VALUES
			(
				?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,
				?,?,?,?,?,?
			)
		`)

		if err != nil {
			customlog.HandleError("Prepare insert of NFeDetalhe", err)
			return err
		}

		_, err = stmt.Exec(
			nNF,
			detail.Prod.CProd,
			detail.Prod.CEAN,
			detail.Prod.XProd,
			detail.Prod.NCM,
			detail.Prod.CEST,
			detail.Prod.CFOP,
			detail.Prod.UCom,
			detail.Prod.QCom,
			detail.Prod.VUnCom,
			detail.Prod.VProd,
			detail.Prod.CEANTrib,
			detail.Prod.UTrib,
			detail.Prod.QTrib,
			detail.Prod.VUnTrib,
			detail.Prod.IndTot,
			hasXPed(detail.Prod.XPed),
			printValue(detail.Imposto.ICMS.ICMS00.Orig, detail.Imposto.ICMS.ICMS10.Orig),
			printValue(detail.Imposto.ICMS.ICMS00.CST, detail.Imposto.ICMS.ICMS10.CST),
			printValue(detail.Imposto.ICMS.ICMS00.ModBC, detail.Imposto.ICMS.ICMS10.ModBC),
			printValue(detail.Imposto.ICMS.ICMS00.VBC, detail.Imposto.ICMS.ICMS10.VBC),
			printValue(detail.Imposto.ICMS.ICMS00.PICMS, detail.Imposto.ICMS.ICMS10.PICMS),
			printValue(detail.Imposto.ICMS.ICMS00.VICMS, detail.Imposto.ICMS.ICMS10.VICMS),
			fillWithZero(detail.Imposto.ICMS.ICMS10.ModBCST),
			fillWithZero(detail.Imposto.ICMS.ICMS10.PMVAST),
			fillWithZero(detail.Imposto.ICMS.ICMS10.VBCST),
			fillWithZero(detail.Imposto.ICMS.ICMS10.PICMSST),
			fillWithZero(detail.Imposto.ICMS.ICMS10.VICMSST),
			fillWithZero(detail.Imposto.IPI.IPITrib.CST),
			fillWithZero(detail.Imposto.IPI.IPITrib.VBC),
			fillWithZero(detail.Imposto.IPI.IPITrib.PIPI),
			fillWithZero(detail.Imposto.IPI.IPITrib.VIPI),
			fillWithZero(detail.Imposto.PIS.PISAliq.CST),
			fillWithZero(detail.Imposto.PIS.PISAliq.VBC),
			fillWithZero(detail.Imposto.PIS.PISAliq.PPIS),
			fillWithZero(detail.Imposto.PIS.PISAliq.VPIS),
			fillWithZero(detail.Imposto.COFINS.COFINSAliq.CST),
			fillWithZero(detail.Imposto.COFINS.COFINSAliq.VBC),
			fillWithZero(detail.Imposto.COFINS.COFINSAliq.PCOFINS),
			fillWithZero(detail.Imposto.COFINS.COFINSAliq.VCOFINS),
			detail.InfAdProd)

		if err != nil {
			customlog.HandleError("Inserting NFeDetalhe", err)
			return err
		}

	}

	return nil
}

func printValue(param1 string, param2 string) string {
	content := param1

	if len(content) == 0 {
		content = param2
	}

	return content
}

func hasXPed(xPed string) string {
	if len(xPed) == 0 {
		return "*** sem ordem de compra ***"
	}

	return xPed
}

func fillWithZero(param string) string {
	content := param

	if len(param) == 0 {
		content = "0"
	}

	return content
}
