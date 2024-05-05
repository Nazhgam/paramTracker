package controller

import (
	"net/http"
	"strings"

	"github.com/Nazhgam/paramTracker/internal/models"
	"github.com/labstack/echo/v4"
)

// @Summary Create a new point
// @Description Create a new point with the provided details
// @Tags Points
// @Accept json
// @Produce json
// @Param req body models.Point true "Point details"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} string "Bad Request"
// @Router /api/v1/point [post]
func (h *handler) CreatePoint(e echo.Context) error {
	var req models.Point

	err := e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	err = h.srv.CreatePoint(e.Request().Context(), req)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusNoContent, nil)
}

// @Summary Get the last ID by alias
// @Description Retrieve the last ID associated with the provided alias
// @Tags Parameters
// @Accept json
// @Produce json
// @Param alias query string true "Alias for the parameter"
// @Success 200 {object} models.IDResponse "Last ID"
// @Failure 400 {object} string "Bad Request"
// @Router /api/v1/lastID [get]
func (h *handler) GetLastID(e echo.Context) error {
	alias := e.QueryParam("alias")
	if strings.TrimSpace(alias) == "" {
		h.lg.Println("error empty alias")
		return e.JSON(http.StatusBadRequest, nil)
	}

	res, err := h.srv.GetLastIDByAlias(e.Request().Context(), alias)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, res)
}
