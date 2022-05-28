package http

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type Server struct {
	svr http.Server
	err error
}

func NewServer(addr string, handle http.Handler) *Server {
	s := new(Server)
	s.err = s.genServer(addr, handle)
	return s
}

func (s *Server) genServer(addr string, handle http.Handler) error {

	svr := http.Server{
		Addr:    addr,
		Handler: handle,
	}

	s.svr = svr
	return nil
}

func (s *Server) Start(ctx context.Context) error {
	err := s.svr.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("HTTP Server Stop")
	return s.svr.Shutdown(ctx)
}
