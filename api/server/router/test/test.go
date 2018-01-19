package test

import (
	"storm/api/server/router"
)

// testRouter is just test
type testRouter struct {
	routes  []router.Route
}

// NewRouter initializes a new test router
func NewRouter() router.Router {
	r := &testRouter{
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the test
func (r *testRouter) Routes() []router.Route {
	return r.routes
}

func (r *testRouter) initRoutes() {
	r.routes = []router.Route{
		// GET
		router.NewGetRoute("/testget/{name:.*}", r.testGet),

		// POST
		router.NewPostRoute("/testpost", r.testPost),

		// DELETE
		router.NewDeleteRoute("/testdelete", r.testDelete),
	}
}