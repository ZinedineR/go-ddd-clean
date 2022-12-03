package api

import (
	"github.com/labstack/echo/v4"
	_healthHandler "latihan/domain/health/handler"
)

// handlerInitialization
var (
	healthHandler *_healthHandler.HTTPHandler
)

func setupRepoService() {
	healthHandler = _healthHandler.NewHTTPHandler()
}

func SetupRouter() *echo.Echo {
	setupRepoService()
	route := echo.New()
	route.GET("/", healthHandler.Index)

	return route
}
