package v2

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/identbase/identserv/pkg/account"
	"github.com/identbase/serv/pkg/server"
)

/*
AuthContext extends echo.Context to include authentication information. */
type AuthContext struct {
	echo.Context
	User account.User
}

/*
NewAuthContext creates an echo.Context object that also contains authentication
information. */
func NewAuthContext(c echo.Context) *AuthContext {
	return &AuthContext{c, account.User{}}
}

/*
TokenFromRequest pulls out the access_token from the http.Request */
func TokenFromRequest(r *http.Request) string {
	t := ""

	a := r.Header.Get("Authorization")
	if a != "" && strings.HasPrefix(a, "Bearer ") {
		t = strings.Split(a, "Bearer ")[1]
	}

	if t != "" {
		return t
	}

	if q, err := url.ParseQuery(r.URL.RawQuery); err == nil {
		t = q.Get("access_token")
	}

	return t
}

/*
AuthRequired middleware indicates that an endpoint requires authentication. */
func (v *V2) AuthRequired(n echo.HandlerFunc) echo.HandlerFunc {
	db, err := v.Context.GetDatabase()
	if err != nil {
		return func(c echo.Context) error {
			c.Logger().Error("panic: ", err)

			return c.JSON(http.StatusInternalServerError, *server.Errors[http.StatusInternalServerError])
		}
	}

	return func(c echo.Context) error {
		cc := NewAuthContext(c)

		if t := TokenFromRequest(c.Request()); t != "" {
			cc.Logger().Infof("Looking up token '%s'", t)
			if a := db.Lookup(t); a != nil {
				u, ok := (*a).(*account.User)
				if !ok {
					cc.Logger().Errorf("panic: store object mismatc, wanted account.User, got %T", a)

					return c.JSON(http.StatusInternalServerError, *server.Errors[http.StatusInternalServerError])
				}

				cc.User = *u

				if err := n(cc); err != nil {
					cc.Error(err)
				}

				return nil
			}

		}

		return c.JSON(http.StatusUnauthorized, *server.Errors[http.StatusUnauthorized])
	}
}
