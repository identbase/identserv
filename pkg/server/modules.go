package server

import (
	"github.com/labstack/echo/v4"
)

// TODO: These feel weird living in the root directory and should probably be
// moved.
/*
Module is essentially a wrapper over the echo Group type. */
type Module struct {
	echo.Group
	Router
}

/*
NewModule generates the underlying echo Group and register the routes. */
func (s *Server) NewModule(p string, r Router) *Module {
	g := s.engine.Group(p)
	rs := r.Routes()

	for i := 0; i < len(rs); i += 1 {
		g.Add(rs[i].Method, rs[i].Path, rs[i].Handler)
	}

	return &Module{
		Group:  *g,
		Router: r,
	}
}
