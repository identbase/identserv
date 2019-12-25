package v1

import (
	"github.com/mihok/identbase/pkg/server"
	"github.com/mihok/identbase/pkg/store"
)

/*
V1 implementation for the routes */
type V1 struct {
	Context
}

// TODO: Add better support for a generic list of context
type Context interface {
	GetDatabase() (*store.InMemory, error)
}

/*
Routes provides a list of routes that this Router will answer to. */
func (v *V1) Routes(c Context) []*server.Route {
	// TODO: Maybe dont push stuff to V1 here?
	v.Context = c

	return []*server.Route{
		// Status check
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/",
				Name:   "Status check",
				// Default: true,
			},
			Handler: v.GetStatus,
		},

		// Key related routes
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/pubkey/:key",
				Name:   "Get key",
			},
			Handler: v.GetKey,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/pubkey/isvalid",
				Name:   "Get key",
			},
			Handler: v.GetKeyValidity,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/pubkey/emphemeral/isvalid",
				Name:   "Get key",
			},
			Handler: v.GetEmphemeralKeyValidity,
		},

		// Lookup routes
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/lookup",
				Name:   "Get lookup",
			},
			Handler: v.GetLookup,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "POST",
				Path:   "/bulk_lookup",
				Name:   "Post bulk lookup",
			},
			Handler: v.PostBulkLookup,
		},
	}
}
