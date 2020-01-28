package v2

import (
	"fmt"

	"github.com/identbase/identserv/pkg/store"
	"github.com/identbase/serv/pkg/server"
)

/*
V2 implementation for the routes */
type V2 struct {
	Context
}

/*
Context provides any state context for V1 routes. */
// TODO: Add better support for a generic list of context
type Context interface {
	GetDatabase() (*store.InMemory, error)
}

/*
Routes provides a list of routes that this Router will answer to. */
func Routes(c Context) []*server.Route {
	v := V2{}

	// TODO: Maybe dont push stuff to V1 here?
	v.Context = c
	pre := "/v2"

	return []*server.Route{
		// Status check
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   fmt.Sprintf("%s", pre),
				Name:   "Status check",
				// Default: true,
			},
			Handler: v.GetStatus,
		},

		// Account routes
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "POST",
				Path:   fmt.Sprintf("%s/account/register", pre),
				Name:   "Account registration",
			},
			Handler: v.PostAccountRegister,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "POST",
				Path:   fmt.Sprintf("%s/account/logout", pre),
				Name:   "Account logout",
			},
			Handler: v.AuthRequired(v.PostAccountLogout),
		},
	}

}
