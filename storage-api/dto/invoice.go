package dto

type InvoiceProductDto struct {
	InvoiceID uint `json:"invoiceId"`
	ProductID uint `json:"productId"`
	Quantity  int  `json:"quantity"`
}
