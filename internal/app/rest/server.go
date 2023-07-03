package rest

import (
	"golang.org/x/net/context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}
func (s *Server) ShuttingDown(ctx context.Context) error {
	log.Println("App is shut down")
	return s.httpServer.Shutdown(ctx)
}
