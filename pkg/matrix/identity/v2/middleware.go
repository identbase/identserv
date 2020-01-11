package v2

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

/*
TokenFromRequest pulls out the access_token from the http.Request */
func TokenFromRequest(r http.Request) string {
	t := ""

	a := r.Header.Get("Authorization")
	if a == "" && strings.HasPrefix(a, "Bearer ") {
		t = strings.Split(a, "Bearer ")[1]
	}

	q, err := url.ParseQuery(r.URL.RawQuery)
	if t == "" && err == nil {
		t = q.Get("access_token")
	}

	return t
}

/*
AuthRequired middleware indicates that an endpoint requires authentication. */
func AuthRequired(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return f(c)
	}
}
