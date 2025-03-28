package dto

type InvoiceProductDto struct {
	InvoiceID uint   `json:"invoiceId"`
	ProductID uint   `json:"productId"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
}
