package management

import (
"storm/api/server/router"
)

// testRouter is just test
type managementRouter struct {
	routes  []router.Route
}

// NewRouter initializes a new test router
func NewRouter() router.Router {
	r := &managementRouter{
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the test
func (r *managementRouter) Routes() []router.Route {
	return r.routes
}

func (r *managementRouter) initRoutes() {
	r.routes = []router.Route{
		// post
		router.NewPostRoute("/management/createcluster", r.createCluster),
	}
}