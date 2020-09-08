package server

import (
	"time"
	"net/http"
	
	"github.com/kangbojk/library_api/api/router"
	"github.com/kangbojk/library_api/entity"
	"github.com/kangbojk/library_api/config"
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