package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/labstack/gommon/log"
)

/*
Config holds all the configuration that is adjustable. */
type Config struct {
	HideBanner bool
	HidePort   bool
	Debug      bool
}

/*
Context is essentially a wrapper for echo's context. */
type Context struct {
	echo.Context
}

/*
Middleware is a wrapper around echo's MiddlewareFunc function. */
// type Middleware func(Context) error

// TODO: Wrap the MiddlewareFunc better.
/*
Use is a wrapper around echo's Use function. */
func (s *Server) Use(m ...echo.MiddlewareFunc) {
	s.engine.Use(m...)
}

/*
Load adds a module into the server to expose */
func (s *Server) Load(m *Module) {
	s.Modules = append(s.Modules, m)
}

/*
Logger returns the (echo) engine logger. */
func (s *Server) Logger() echo.Logger {
	return s.engine.Logger
}

/*
Start begins serving traffic. */
func (s *Server) Start(h *http.Server) error {
	return s.engine.StartServer(h)
}

/*
Server is the core process. */
type Server struct {
	engine  *echo.Echo
	Modules []*Module
}

/*
New creates a new instance of Server based on a Config passed to it. */
func New(c Config) *Server {
	e := echo.New()

	if c.Debug {
		e.Logger.SetLevel(log.DEBUG)
	} else {
		e.Logger.SetLevel(log.INFO)
	}

	e.HideBanner = c.HideBanner
	e.HidePort = c.HidePort

	e.HTTPErrorHandler = func(e error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := e.(*echo.HTTPError); ok {
			code = he.Code
		}

		if err := c.JSON(code, Errors[code]); err != nil {
			c.Logger().Error(err)
		}

		c.Logger().Error(e)
	}

	return &Server{
		engine: e,
	}
}
