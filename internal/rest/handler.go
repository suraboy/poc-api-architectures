package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler interface {
	GetUser(c echo.Context) error
}

type handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return &handler{svc}
}

func (hdl *handler) GetUser(c echo.Context) error {
	response, err := hdl.svc.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}
