package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kangbojk/library_api/api"
	"github.com/kangbojk/library_api/config"
	"github.com/kangbojk/library_api/entity"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Println("---------")
		log.Println(err.Error())
	}
	log.Println("connecting to db ", dataSourceName)
	defer db.Close()

	bookRepo := book.NewMysqlRepository(db)
	srv := server.NewServer(bookRepo)
	log.Fatal(srv.ListenAndServe())
}
