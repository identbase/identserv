package v1

import (
	"github.com/identbase/identserv/pkg/store"
	"github.com/identbase/serv/pkg/server"
)

/*
V1 implementation for the routes */
type V1 struct {
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
// TODO: This is essentially an init func for V1 class, maybe it shouldnt be
func Routes(c Context) []*server.Route {
	v := V1{}

	// TODO: Maybe dont push stuff to V1 here?
	v.Context = c
	pre := "/v1"

	return []*server.Route{
		// Status check
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   pre + "/",
				Name:   "Status check",
				// Default: true,
			},
			Handler: v.GetStatus,
		},

		// Key related routes
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   pre + "/pubkey/:key",
				Name:   "Get key",
			},
			Handler: v.GetKey,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   pre + "/pubkey/isvalid",
				Name:   "Get key",
			},
			Handler: v.GetKeyValidity,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   pre + "/pubkey/emphemeral/isvalid",
				Name:   "Get key",
			},
			Handler: v.GetEmphemeralKeyValidity,
		},

		// Lookup routes
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   pre + "/lookup",
				Name:   "Get lookup",
			},
			Handler: v.GetLookup,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "POST",
				Path:   pre + "/bulk_lookup",
				Name:   "Post bulk lookup",
			},
			Handler: v.PostBulkLookup,
		},
	}
}
