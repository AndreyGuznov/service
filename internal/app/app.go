package app

import (
	"serv/config"
	"serv/internal/controller"
	"serv/internal/repo"
	"serv/internal/service"
	"serv/pkg/cache"
	"serv/pkg/httpserver"
	"serv/pkg/logger"
	"serv/pkg/reindex"
	"time"
)

func Run() {

	config.InitConfig()

	reindex.Conn()

	defer reindex.Conn().Close()

	err := repo.StartNewspaces("city")

	if err != nil {
		logger.Err("Failed work with NamespaceCity", err)
	}

	cache.InitCache()

	ticker := time.NewTicker(time.Duration(config.Instance.Time) * time.Minute)

	if err := service.RefresfData(); err != nil {
		logger.Err("Error of updating cach", err)
	}

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				err := service.RefresfData()
				if err != nil {
					logger.Err("Error of updating cach", err)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	controller := controller.NewController()

	controller.GetRoutes()

	handler := httpserver.NewHTTPHandler(controller)

	httpServer := httpserver.NewHTTPRestServer(":"+config.Instance.Server.Addr, handler)

	err = httpServer.Serve()
	if err != nil {
		logger.Err("Error of Server", err)
	}

}
