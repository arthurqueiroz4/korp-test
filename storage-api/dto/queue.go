package dto

type MessageForBilling struct {
	InvoiceID uint   `json:"invoiceId"`
	Detail    string `json:"detail"`
	Status    string `json:"status"`
}
