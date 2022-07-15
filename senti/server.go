package senti

import (
	"context"
	"net/http"
	"time"
)

//Server base struct http server
type Server struct {
	httpVeryCuteServer *http.Server
}

//RunMyServerApi start http server
func (s *Server) RunMyServerApi(port string, handler http.Handler) error {
	s.httpVeryCuteServer = &http.Server{
		Addr:           "127.0.0.1:" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 21, // 2 MB
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
	}
	return s.httpVeryCuteServer.ListenAndServe()
}

//Shutdown stop http server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpVeryCuteServer.Shutdown(ctx)
}
