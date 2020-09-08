package book

import (
	"time"

	"github.com/google/uuid"
)

func NewFixtureBook() *Book {
	return &Book{
		ID:        uuid.New(),
		Title:     "How to Be Rich",
		Author:    "J. Paul Getty",
		Pages:     224,
		Quantity:  1,
		CreatedAt: time.Now(),
	}
}
