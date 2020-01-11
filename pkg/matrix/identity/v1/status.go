package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
Status is a simple JSON response object. */
// TODO: Add process metrics to monitor the server
type Status struct {
	// Message string `json:"message"`
}

/*
GetStatus is used for auto-discovery and health checks. */
func (v *V1) GetStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, Status{})
}
