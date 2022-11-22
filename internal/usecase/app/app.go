package app

import (
	"serv/internal/controller"
	"serv/internal/usecase/repo"
	"serv/pkg/httpserver"
	"serv/pkg/logger"
	"serv/pkg/reindex"
)

func Run() {

	reindex.Conn()

	defer reindex.Conn().Close()

	err := repo.StartNewspaces("city", "carbase")

	if err != nil {
		logger.Err("Failed work with NamespaceCity", err)
	}

	controller := controller.NewController()
	controller.GetRoutes()

	handler := httpserver.NewHTTPHandler(controller)

	httpServer := httpserver.NewHTTPRestServer(":"+"1234", handler)

	err = httpServer.Serve()
	if err != nil {
		logger.Err("Error of Server", err)
	}

}
