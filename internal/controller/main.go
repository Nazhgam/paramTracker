package controller

import (
	"log"

	"github.com/Nazhgam/paramTracker/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	CreateGps(e echo.Context) error
	GetAllGps(e echo.Context) error
	CreateParamStr(e echo.Context) error
	CreateParamInt(e echo.Context) error
	CreatePoint(e echo.Context) error
	GetLastID(e echo.Context) error
}

type handler struct {
	lg  *log.Logger
	srv service.Service
}

func New(lg *log.Logger, srv service.Service) Handler {
	return &handler{
		lg:  lg,
		srv: srv,
	}
}
