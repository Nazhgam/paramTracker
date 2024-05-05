package controller

import (
	"net/http"
	"strings"

	"github.com/Nazhgam/paramTracker/internal/models"
	"github.com/labstack/echo/v4"
)

// @Summary Create a new string parameter
// @Description Create a new string parameter with the provided alias and value
// @Tags Parameters
// @Accept json
// @Produce json
// @Param alias query string true "Alias for the parameter"
// @Param req body models.ParamStr true "Parameter details"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} string "Bad Request"
// @Router /api/v1/param-str-by-alias [post]
func (h *handler) CreateParamStr(e echo.Context) error {
	var req models.ParamStr

	if err := e.Bind(&req); err != nil {
		h.lg.Printf("error bind req in create param_str: %s", err)
		return e.JSON(http.StatusBadRequest, err)
	}

	alias := e.QueryParam("alias")
	if strings.TrimSpace(alias) == "" {
		h.lg.Println("error empty alias")
		return e.JSON(http.StatusBadRequest, nil)
	}

	err := h.srv.CreateParamStr(e.Request().Context(), alias, req)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusNoContent, nil)
}

// @Summary Create a new integer parameter
// @Description Create a new integer parameter with the provided alias and value
// @Tags Parameters
// @Accept json
// @Produce json
// @Param alias query string true "Alias for the parameter"
// @Param req body models.ParamInt true "Parameter details"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} string "Bad Request"
// @Router /api/v1/param-int-by-alias [post]
func (h *handler) CreateParamInt(e echo.Context) error {
	var req models.ParamInt

	if err := e.Bind(&req); err != nil {
		h.lg.Printf("error bind req in create param_int: %s", err)
		return e.JSON(http.StatusBadRequest, err)
	}

	alias := e.QueryParam("alias")
	if strings.TrimSpace(alias) == "" {
		h.lg.Println("error empty alias")
		return e.JSON(http.StatusBadRequest, nil)
	}

	err := h.srv.CreateParamInt(e.Request().Context(), alias, req)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusNoContent, nil)
}
