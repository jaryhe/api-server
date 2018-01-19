package cluster

import (
"storm/api/server/router"
)

// cluster routers
type clusterRouter struct {
	routes  []router.Route
}

// NewRouter initializes a new test router
func NewRouter() router.Router {
	r := &clusterRouter{
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the test
func (r *clusterRouter) Routes() []router.Route {
	return r.routes
}

func (r *clusterRouter) initRoutes() {
	r.routes = []router.Route{
		// post
		router.NewPostRoute("/cluster/create", r.createCluster),
		router.NewPostRoute("/cluster/addstoragenode", r.addStorageNode),
		router.NewPostRoute("/cluster/addmetanode", r.addMetaNode),
		router.NewPostRoute("/cluster/addclientnode", r.addClientNode),
		router.NewPostRoute("/cluster/getstoragestatus", r.getStorageStatus),
		router.NewPostRoute("/cluster/getmetastatus", r.getMetaStatus),
		router.NewPostRoute("/cluster/getclientstatus", r.getClientStatus),
	}
}