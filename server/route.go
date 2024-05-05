package server

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s Server) Route(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")
	v1.POST("/gps-by-alias", s.handler.CreateGps)
	v1.POST("/param-str-by-alias", s.handler.CreateParamStr)
	v1.POST("/param-int-by-alias", s.handler.CreateParamInt)
	v1.POST("/point", s.handler.CreatePoint)
	v1.GET("/lastID", s.handler.GetLastID)
	v1.GET("/gps", s.handler.GetAllGps)
}
