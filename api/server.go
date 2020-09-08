package server

import (
	"net/http"
	"time"

	"github.com/kangbojk/library_api/api/router"
	"github.com/kangbojk/library_api/config"
	"github.com/kangbojk/library_api/entity"
)

func NewServer(repo book.Repository) *http.Server {
	handler := router.NewRouter(repo)

	s := &http.Server{
		Addr:         config.PORT,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s
}
