package identity

import (
	"github.com/identbase/identserv/pkg/matrix/identity/v1"
	"github.com/identbase/identserv/pkg/matrix/identity/v2"
	"github.com/identbase/identserv/pkg/store"
	"github.com/identbase/serv/pkg/server"
)

/*
Matrix implements the Router and Context interface. */
// TODO: This is a shortcut implementation of "Matrix" since we are only
// supporting identity services and not the entire matrix protocol. We should
// consider moving this down the folder tree somewhere to make more sense
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

/*
AddDatabase allows another thing to add a store.InMemory database to use. */
// TODO: Support multiple databases?
func (m *Matrix) AddDatabase(d *store.InMemory) {
	m.Database = d
}

/*
GetDatabase returns the database context. */
func (m *Matrix) GetDatabase() (*store.InMemory, error) {
	return m.Database, nil
}

/*
Routes provides a list of routes that this Router will answer to. */
func (m *Matrix) Routes() []*server.Route {
	// /v1
	r := append(v1.Routes(m))

	// /v2
	r = append(v2.Routes(m), r...)

	return r
}
