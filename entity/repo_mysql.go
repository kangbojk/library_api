package book

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

//mysqlRepo mysql repo
type mysqlRepo struct {
	db *sql.DB
}

//NewMysqlRepository create new repository
func NewMysqlRepository(db *sql.DB) *mysqlRepo {
	return &mysqlRepo{
		db: db,
	}
}

//Create a book
func (r *mysqlRepo) Create(b *Book) (uuid.UUID, error) {
	// convert UUID(36 chars) to binary(16)
	stmt, err := r.db.Prepare(`insert into book (id, title, author, pages, quantity, created_at) values(UUID_TO_BIN(?),?,?,?,?,?)`)
	if err != nil {
		return b.ID, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		b.ID,
		b.Title,
		b.Author,
		b.Pages,
		b.Quantity,
		time.Now().Format(layoutISO),
	)
	if err != nil {
		return b.ID, err
	}

	return b.ID, nil
}

//Get a book
func (r *mysqlRepo) Get(id uuid.UUID) (*Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book where id = UUID_TO_BIN(?)`)
	if err != nil {
		return nil, err
	}
	var b Book
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
	}
	return &b, nil
}

//Update a book
func (r *mysqlRepo) Update(b *Book) error {
	b.UpdatedAt = time.Now()
	_, err := r.db.Exec("update book set title = ?, author = ?, pages = ?, quantity = ?, updated_at = ? where id = UUID_TO_BIN(?)", b.Title, b.Author, b.Pages, b.Quantity, b.UpdatedAt.Format(layoutISO), b.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search books
func (r *mysqlRepo) Search(query string) ([]*Book, error) {
	stmt, err := r.db.Prepare(`select BIN_TO_UUID(id) AS uuid, title, author, pages, quantity, created_at from book where title like ?`)
	if err != nil {
		return nil, err
	}
	var books []*Book
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	if len(books) == 0 {
		return nil, ErrNotFound
	}
	return books, nil
}

//List books
func (r *mysqlRepo) List() ([]*Book, error) {
	stmt, err := r.db.Prepare(`select BIN_TO_UUID(id) AS uuid, title, author, pages, quantity, created_at from book`)
	if err != nil {
		return nil, err
	}
	var books []*Book
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	if len(books) == 0 {
		return nil, ErrNotFound
	}
	return books, nil
}

//Delete a book
func (r *mysqlRepo) Delete(id uuid.UUID) error {
	_, err := r.db.Exec("delete from book where id = UUID_TO_BIN(?)", id)
	if err != nil {
		return err
	}
	return nil
}
