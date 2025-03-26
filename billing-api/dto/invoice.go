package dto

type InvoiceCreateDto struct {
	Numeration string              `json:"numeration"`
	Products   []InvoiceProductDto `json:"products"`
}

type InvoiceReadDto struct {
	ID         uint                `json:"id"`
	Status     string              `json:"status"`
	Numeration string              `json:"numeration"`
	Items      []InvoiceProductDto `json:"products"`
}

type InvoiceProductDto struct {
	ID       uint `json:"id"`
	Quantity int  `json:"quantity"`
}
