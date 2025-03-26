package dto

type MessageForBilling struct {
	InvoiceID uint   `json:"invoiceId"`
	Status    string `json:"status"`
}
