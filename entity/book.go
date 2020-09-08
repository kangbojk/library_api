package book

import (
	"time"
	"errors"

	"github.com/google/uuid"
)

//Book data
type Book struct {
	ID        uuid.UUID
	Title     string
	Author    string
	Pages     int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

//repository interface
type Repository interface {
	Get(id uuid.UUID) (*Book, error)
	Search(query string) ([]*Book, error)
	List() ([]*Book, error)

	Create(b *Book) (uuid.UUID, error)
	Update(b *Book) error
	Delete(id uuid.UUID) error
}

var ErrNotFound = errors.New("Not found\n")