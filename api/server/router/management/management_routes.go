package management

import (
	"golang.org/x/net/context"
	"net/http"
	"storm/api/server/httputils"
	"strings"
)

func (t *managementRouter)createCluster(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}
	nodestr := r.Form["nodelist"][0]
	clusterip := r.Form["clusterip"][0]
    nodelist := strings.Split(nodestr,",")
	cluster := &Cluster{
		ClusterIP:clusterip,
		NodeList:nodelist,
	}
	err := cluster.InitCluster()
	if err != nil{

	}
	return httputils.WriteJSON(w, http.StatusCreated,"test post interface")
}