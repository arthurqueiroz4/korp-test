package dto

type Page[T any] struct {
	Content []T `json:"content"`
	Page    int `json:"page"`
	Size    int `json:"size"`
}
