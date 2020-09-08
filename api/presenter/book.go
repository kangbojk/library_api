package presenter

import (
	"github.com/google/uuid"
)

//Book data
type Book struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Pages    int       `json:"pages"`
	Quantity int       `json:"quantity"`
}
