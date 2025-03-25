package dto

type ProductCreateDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Balance     int    `json:"balance"`
}

type ProductReadDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Balance     int    `json:"balance"`
}
