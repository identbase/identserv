package server

import (
	"github.com/labstack/echo/v4"
)

/*
RouteMeta is a clone of the echo.Route object that we use to extend our
routes with. */
type RouteMeta struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Name   string `json:"name"`
}

/*
Route defines a Handler and some metadata about the route. */
type Route struct {
	RouteMeta
	Handler echo.HandlerFunc
}

/*
Router is a group of routes to expose and serve through the server. */
type Router interface {
	Routes() []*Route
	AddContext(key string, value interface{})
	GetContext(key string) interface{}
}
