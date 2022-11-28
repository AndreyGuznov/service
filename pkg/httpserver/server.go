package httpserver

import (
	"net/http"
	"serv/pkg/logger"
)

type HTTPRestServer struct {
	address string
	wrapped *http.Server
}

func NewHTTPRestServer(address string, handler http.Handler) *HTTPRestServer {
	httpsrv := http.Server{
		Addr:    address,
		Handler: handler,
	}

	return &HTTPRestServer{wrapped: &httpsrv, address: address}
}

func (server *HTTPRestServer) Serve() error {
	logger.Info("Serving on " + server.address)
	err := server.wrapped.ListenAndServe()

	if err != http.ErrServerClosed {
		logger.Err("Server crashed", err)
	}

	return err
}
