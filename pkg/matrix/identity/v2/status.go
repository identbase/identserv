package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: Add process metrics to monitor the server
/*
Status is a simple JSON response object. */
type Status struct {
	// Message string `json:"message"`
}

/*
GetStatus is used for auto-discovery and health checks. */
func (v *V2) GetStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, Status{})
}
