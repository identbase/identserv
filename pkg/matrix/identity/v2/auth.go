package v2

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/identbase/getting"
	"github.com/identbase/getting/pkg/resource/representor"
	"github.com/identbase/identserv/pkg/account"
	"github.com/identbase/serv/pkg/server"
)

/*
Account interface represents an account object held by the server to connect
a user from a matrix server to its threepids. */
type Account interface {
	GenerateToken()
	Save() error
}

/*
RegisterRequest is the JSON request expected by PostAccountRegister. */
type RegisterRequest struct {
	AccessToken string `json:"access_token" form:"access_token" query:"access_token"`
	Type        string `json:"token_type" form:"token_type" query:"token_type"`
	Domain      string `json:"matrix_server_name" form:"matrix_server_name" query:"matrix_server_name"`
	Expires     int    `json:"expires_in" form:"expires_in" query:"expires_in"`
}

/*
RegisterFederationResponse is the JSON response from a Matrix server. */
type RegisterFederationResponse struct {
	representor.HALBody
	// Id           string `json:"sub"`
	// ErrorCode    string `json:"errcode"`
	// ErrorMessage string `json:"error"`
}

/*
RegisterResponse is the JSON response sent by PostAccountRegister. */
type RegisterResponse struct {
	representor.HALBody
}

/*
NewRegisterResponse creates a RegisterResponse object to respond with. */
func NewRegisterResponse(t string) *RegisterResponse {
	r := RegisterResponse{
		representor.HALBody{
			Links: map[string][]representor.HALLink{
				"self": []representor.HALLink{
					representor.HALLink{
						HRef:  "/_matrix/identity/v2/account/register",
						Title: "Register",
					},
				},
				"base": []representor.HALLink{
					representor.HALLink{
						HRef:  "/_matrix/identity/v2/",
						Title: "Default",
					},
				},
				"account": []representor.HALLink{
					representor.HALLink{
						HRef:  "/_matrix/identity/v2/account",
						Title: "Account information",
					},
				},
				"logout": []representor.HALLink{
					representor.HALLink{
						HRef:  "/_matrix/identity/v2/account/logout",
						Title: "Logout",
					},
				},
			},
			Properties: map[string]interface{}{
				"token": t,
				// `token` is correct for the spec, but we released with `access_token`
				//  for a substantial amount of time. Serve both to make spec-compliant clients
				//  happy.
				"access_token": t,
			},
		},
	}

	return &r
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
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, *server.NewHALError("E_BAD_JSON", "Malformed JSON"))
	}

	d := fmt.Sprintf("http://%s", req.Domain)
	g, err := getting.New(d)
	if err != nil {
		c.Logger().Error("Unable to set up Getting")
		return c.JSON(http.StatusInternalServerError, *server.Errors[http.StatusInternalServerError])
	}

	f, err := g.Follow("openid", map[string]string{
		"access_token": req.AccessToken,
	})
	if err != nil {
		c.Logger().Errorf("Unable to find links: %v", err)
		return c.JSON(http.StatusInternalServerError, *server.Errors[http.StatusInternalServerError])
	}

	r, err := f.Get()
	if err != nil {
		c.Logger().Errorf("Unable to get resource: %v", err)
		return c.JSON(http.StatusInternalServerError, *server.Errors[http.StatusInternalServerError])
	}

	if rfr, ok := r.(representor.HALBody); ok {
		if sub, ok := rfr.Properties["sub"].(string); ok && sub != "" {
			c.Logger().Infof("Registering new account: '%s'", sub)
			a := account.New(sub)
			if err := a.GenerateToken(); err != nil {
				c.Logger().Errorf("Unable to generate token: %v", err)
				return c.JSON(http.StatusInternalServerError, *server.Errors[http.StatusInternalServerError])
			}

			d, err := v.GetDatabase()
			if err != nil {
				c.Logger().Errorf("Unable to get database: %v", err)
				return c.JSON(http.StatusInternalServerError, *server.Errors[http.StatusInternalServerError])
			}

			d.Put(a)

			return c.JSON(http.StatusOK, NewRegisterResponse(a.Token))
		}

		// TODO: Handle other errors passed by the Federation server?
		return c.JSON(http.StatusUnauthorized, *server.Errors[http.StatusUnauthorized])
	}

	return c.JSON(http.StatusInternalServerError, *server.Errors[http.StatusInternalServerError])
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
