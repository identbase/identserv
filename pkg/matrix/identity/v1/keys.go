package identity

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mihok/identbase/pkg/server"
	"github.com/mihok/identbase/pkg/store"
)

/*
Key is a simple JSON response object that carries a unpadded Base64 public key. */
type Key struct {
	// Message string `json:"message"`
}

/*
GetKey is used to get the public key passed by :key

GET /v1/pubkey/:id

Parameters:

:id - Required. The ID of the key. This should take the form
algorithm:identifier where algorithm identifies the signing algorithm, and the
identifier is an opaque string.

Reference:
https://matrix.org/docs/spec/identity_service/r0.2.1#id12 */
func (m *Matrix) GetKey(c echo.Context) error {
	db := m.GetContext("database").(*store.InMemory)

	id := c.Param("id")

	if key := db.Lookup(id); key != nil {
		if err := c.JSON(http.StatusOK, key); err != nil {
			code := http.StatusInternalServerError
			if he, ok := err.(*echo.HTTPError); ok {
				code = he.Code
			}

			c.Logger().Error(err)

			return c.JSON(code, server.Errors[code])
		}
	}

	return c.JSON(http.StatusNotFound, server.Errors[http.StatusNotFound])
}

/*
GetKeyValidity

GET /v1/pubkey/isvalid
GET /v1/pubkey/emphemeral/isvalid

Query Parameters:

References:
https://matrix.org/docs/spec/identity_service/r0.2.1#get-matrix-identity-api-v1-pubkey-isvalid */
func (m *Matrix) GetKeyValidity(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, server.Errors[http.StatusNotImplemented])
}

/*
GetEmphemeralKeyValidity

GET /v1/pubkey/emphemeral/isvalid

Query Parameters:

References:
https://matrix.org/docs/spec/identity_service/r0.2.1#get-matrix-identity-api-v1-pubkey-ephemeral-isvalid */
func (m *Matrix) GetEmphemeralKeyValidity(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, server.Errors[http.StatusNotImplemented])
}
