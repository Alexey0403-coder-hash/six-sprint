package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func New(logger *log.Logger) *Server {
	r := chi.NewRouter()

	r.Get("/", handlers.MainHandler)
	r.Post("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger:     logger,
		httpServer: httpServer,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
