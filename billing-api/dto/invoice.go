package dto

type InvoiceCreateDto struct {
	Numeration string              `json:"numeration"`
	Products   []InvoiceProductDto `json:"products"`
}

type InvoiceReadDto struct {
	ID         uint                `json:"id"`
	Status     string              `json:"status"`
	Numeration string              `json:"numeration"`
	Detail     string              `json:"detail"`
	Items      []InvoiceProductDto `json:"products"`
}

type InvoiceProductDto struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type InvoiceProductRecvDto struct {
	InvoiceId uint   `json:"invoiceId"`
	Status    string `json:"status"`
	Detail    string `json:"detail"`
}
