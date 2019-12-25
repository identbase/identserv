package identity

import (
	"github.com/identbase/identserv/pkg/server"
	"github.com/identbase/identserv/pkg/store"

	"github.com/identbase/identserv/pkg/matrix/v1"
)

// TODO: This is a shortcut implementation of "Matrix" since we are only
// supporting identity services and not the entire matrix protocol. We should
// consider moving this down the folder tree somewhere to make more sense
/*
Matrix implements the Router and Context interface. */
type Matrix struct {
	// TODO: Use a more robust database
	Database *store.InMemory
}

/*
Context is more or less a less-generic context store to use in place of
a generic context store since golang doesnt support generics (yet?). */
type Context interface {
	// Database
	AddDatabase(*store.InMemory)
	GetDatabase() (*store.InMemory, error)
}

// TODO: Support multiple databases?
/*
AddDatabase allows another thing to add a store.InMemory database to use. */
func (m *Matrix) AddDatabase(d *store.InMemory) {
	m.Database = d
}

/*
GetDatabase returns the database context. */
func (m *Matrix) GetDatabase() (error, *store.InMemory) {
	return m.Database
}

/*
Routes provides a list of routes that this Router will answer to. */
func (m *Matrix) Routes() []*server.Route {
	r := append(v1.Routes(context))

	return []*server.Route{
		// Status check
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/v1",
				Name:   "Status check",
			},
			Handler: m.GetStatus,
		},

		// Key related routes
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/v1/pubkey/:key",
				Name:   "Get key",
			},
			Handler: m.GetKey,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/v1/pubkey/isvalid",
				Name:   "Get key",
			},
			Handler: m.GetKeyValidity,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/v1/pubkey/emphemeral/isvalid",
				Name:   "Get key",
			},
			Handler: m.GetEmphemeralKeyValidity,
		},

		// Lookup routes
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "GET",
				Path:   "/v1/lookup",
				Name:   "Get lookup",
			},
			Handler: m.GetLookup,
		},
		&server.Route{
			RouteMeta: server.RouteMeta{
				Method: "POST",
				Path:   "/v1/bulk_lookup",
				Name:   "Post bulk lookup",
			},
			Handler: m.PostBulkLookup,
		},
	}
}

/*
Matrix Error codes. */
const CodeMissingParam = "M_MISSING_PARAM"
const CodeInvalidParam = "M_INVALID_PARAM"
const CodeSessionNotValidated = "M_SESSION_NOT_VALIDATED"
const CodeNoValidSession = "M_NO_VALID_SESSION"
const CodeSessionExpired = "M_SESSION_EXPIRED"
const CodeInvalidEmail = "M_INVALID_EMAIL"
const CodeEmailSendError = "M_EMAIL_SEND_ERROR"
const CodeInvalidAddress = "M_INVALID_ADDRESS"
const CodeSendError = "M_SEND_ERROR"
const CodeUnrecognized = "M_UNRECOGNIZED"
const CodeThreePIDInUse = "M_THREEPID_IN_USE"
const CodeUnknown = "M_UNKNOWN"
