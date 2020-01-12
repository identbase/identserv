package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/identbase/serv/pkg/server"
)

/*
AccountRegisterRequest is the JSON request expected by PostAccountRegister. */
type AccountRegisterRequest struct {
	AccessToken string `json:"access_token" form:"access_token" query:"access_token"`
	Type        string `json:"token_type" form:"token_type" query:"token_type"`
	Domain      string `json:"matrix_server_name" form:"matrix_server_name" query:"matrix_server_name"`
	Expires     int    `json:"expires_in" form:"expires_in" query:"expires_in"`
}

/*
AccountRegisterResponse is the JSON response sent by PostAccountRegister. */
type AccountRegisterResponse struct {
	Token string `json:"token"`
}

/*
PostAccountRegister registers a new matrix user in the server.

POST /v2/account/register

Post Parameters:
access_token - Required. A string access token the consumer may use to verify
  the identity of the person who generated the token.
token_type - Required. A string token type ("Bearer").
matrix_server_name - Required. A string homeserver domain the consumer should
  use when attempting to verify the user's identity.
expires_in - Required. An integer of seconds before this token expires and a
  new one must be generated.

Reference:
https://matrix.org/docs/spec/identity_service/r0.3.0#post-matrix-identity-v2-account-register */
func (v *V2) PostAccountRegister(c echo.Context) error {
	req := new(AccountRegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, *server.NewHALError("E_BAD_JSON", "Malformed JSON"))
	}

	return c.JSON(http.StatusNotImplemented, *server.Errors[http.StatusNotImplemented])
}

/*
GetAccount gets an account based on :user_id.

GET /v2/account

Query Parameters:

user_id - Required. A string of the user ID which is registered.

Reference:
https://matrix.org/docs/spec/identity_service/r0.3.0#get-matrix-identity-v2-account */
func (v *V2) GetAccount(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, server.Errors[http.StatusNotImplemented])
}

/*
PostAccountLogout logs a user out of the server.

POST /v2/account/logout

Reference:
https://matrix.org/docs/spec/identity_service/r0.3.0#post-matrix-identity-v2-account-logout */
func (v *V2) PostAccountLogout(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, server.Errors[http.StatusNotImplemented])
}
