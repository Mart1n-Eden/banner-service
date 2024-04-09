package server

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

// Инициализация сервера
func NewServer(h http.Handler, addr string) *Server {

	s := &Server{
		// конфигурация http сервера
		httpServer: &http.Server{
			Addr:    ":" + addr,
			Handler: h,
			//
		},
	}

	return s
}

// Запуск сервера
func (s *Server) Run() {
	// ...
	s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
