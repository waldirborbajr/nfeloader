package handlers

import "github.com/waldirborbajr/nfeloader/internal/services"

type QuotationHandler struct {
	symbol            string
	regularMarketTime float32
	QuotationService  services.QuotationService
}

func NewQuotation(quotationService *services.QuotationService) *QuotationHandler {
	return &QuotationHandler{
		QuotationService: *quotationService,
	}
}

func (quotationHandler QuotationHandler) GetQuote(symbol string) (*float32, error) {
	quotation, err := quotationHandler.QuotationService.GetQuote(symbol)
	if err != nil {
		return nil, err
	}
	return quotation, nil
}
