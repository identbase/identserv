package identity

import (
	"github.com/mihok/identbase/pkg/server"
	"github.com/mihok/identbase/pkg/store"
)

// TODO: This is a shortcut implementation of "Matrix" since we are only
// supporting identity services and not the entire matrix protocol. We should
// consider moving this down the folder tree somewhere to make more sense
/*
Matrix implements the Router interface. */
type Matrix struct {
	Database *store.InMemory
}

func (m *Matrix) AddContext(k string, v interface{}) {
	// TODO: Throw an error or warning when we're trying to add context
	// that it dosent support
	if k == "database" {
		m.Database = v.(*store.InMemory)
	}
}

func (m *Matrix) GetContext(k string) interface{} {
	if k == "database" {
		return m.Database
	}

	// TODO: Throw an error or warning when we're trying to grab a context
	// that it dosent support
	return nil
}

/*
Routes provides a list of routes that this Router will answer to. */
func (m *Matrix) Routes() []*server.Route {
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
Matrix Error code. */
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
