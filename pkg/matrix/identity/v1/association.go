package identity

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mihok/identbase/pkg/server"
	// "github.com/mihok/identbase/pkg/store"
)

/*
GetLookup

GET /v1/lookup

Reference:
https://matrix.org/docs/spec/identity_service/r0.2.1#get-matrix-identity-api-v1-lookup */
func (m *Matrix) GetLookup(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, server.Errors[http.StatusNotImplemented])
}

/*
PostBulkLookup

POST /v1/bulk_lookup

Post Parameters:

threepids - Required. An array of tuples comprised of 3PID type and the address
to look up.

Reference:
https://matrix.org/docs/spec/identity_service/r0.2.1#post-matrix-identity-api-v1-bulk-lookup */
func (m *Matrix) PostBulkLookup(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, server.Errors[http.StatusNotImplemented])
}
