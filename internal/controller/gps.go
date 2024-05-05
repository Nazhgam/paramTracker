package controller

import (
	"net/http"
	"strings"

	"github.com/Nazhgam/paramTracker/internal/models"
	"github.com/labstack/echo/v4"
)

// @Summary Create GPS record
// @Description Create a new GPS record
// @Tags GPS
// @Accept json
// @Produce json
// @Param alias query string true "Alias for GPS record"
// @Param body body models.GPS true "GPS record object"
// @Success 204 {string} string "No content"
// @Failure 400 {string} string "Bad request"
// @Router /api/v1//gps-by-alias [post]
func (h *handler) CreateGps(e echo.Context) error {
	var req models.GPS

	if err := e.Bind(&req); err != nil {
		h.lg.Printf("error bind req in create gps: %s", err)
		return e.JSON(http.StatusBadRequest, err)
	}

	alias := e.QueryParam("alias")
	if strings.TrimSpace(alias) == "" {
		h.lg.Println("error empty alias")
		return e.JSON(http.StatusBadRequest, nil)
	}

	err := h.srv.CreateGps(e.Request().Context(), alias, req)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusNoContent, nil)
}

// @Summary Get all GPS records
// @Description Get all GPS records
// @Tags GPS
// @Accept json
// @Produce json
// @Success 200 {array} models.GPS "List of GPS records"
// @Router /gps [get]
func (h *handler) GetAllGps(e echo.Context) error {
	res, err := h.srv.GetAllGps(e.Request().Context())
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, res)

}
