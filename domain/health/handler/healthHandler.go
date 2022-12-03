package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HTTPHandler struct {
	// add service if needed
}

func NewHTTPHandler() *HTTPHandler {
	// return function with service
	return &HTTPHandler{}
}

func (h *HTTPHandler) Index(eCtx echo.Context) error {

	return eCtx.JSON(http.StatusOK, "Healthy")
}
