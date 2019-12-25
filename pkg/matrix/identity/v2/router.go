package v2

import (
	"github.com/mihok/identbase/pkg/server"
	"github.com/mihok/identbase/pkg/store"
)

/*
V1 implementation for the routes */
type V2 struct {
	Context
}

// TODO: Add better support for a generic list of context
type Context interface {
	GetDatabase() (*store.InMemory, error)
}

/*
Routes provides a list of routes that this Router will answer to. */
func (v *V2) Routes(c Context) []*server.Route {
	// TODO: Maybe dont push stuff to V1 here?
	v.Context = c
	pre := "/v2"

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
