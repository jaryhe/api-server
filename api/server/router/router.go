package router

import "storm/api/server/httputils"

// Router defines an interface to specify a group of routes to add to the storm server.
type Router interface {
	// Routes returns the list of routes to add to the storm server.
	Routes() []Route
}

// Route defines an individual API route in the storm server.
type Route interface {
	// Handler returns the raw function to create the http handler.
	Handler() httputils.APIFunc
	// Method returns the http method that the route responds to.
	Method() string
	// Path returns the subpath where the route responds to.
	Path() string
}
