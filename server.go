package tc

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	//httpServer *http.Server
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, //1 MB
		Handler:        handler,
		ReadTimeout:    600 * time.Second,
		WriteTimeout:   600 * time.Second,
	}
	//TODO ХАРДКОД для боевого сервер на https
	
        //return s.httpServer.ListenAndServeTLS("/etc/ssl/jusanmobile.kz/jusanmobile.kz.pem","/etc/ssl/jusanmobile.kz/jusanmobile.kz.key")
        return s.httpServer.ListenAndServeTLS("/etc/ssl/jusanmobile.kz/JM2021.crt","/etc/ssl/jusanmobile.kz/JM2021.key")
	
        //return s.httpServer.ListenAndServeTLS("/etc/ssl/kaztranscom.kz/kaztranscom.kz.pem","/etc/ssl/kaztranscom.kz/kaztranscom.kz.key")
	//return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
