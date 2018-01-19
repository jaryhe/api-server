package cluster
import(
	"golang.org/x/net/context"
	"net/http"
	"storm/api/server/httputils"
	"strings"
	"storm/api/types"
)

func (t *clusterRouter)createCluster(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}
	managenodestr := r.Form["managenodelist"][0]
	clusterip := r.Form["clusterip"][0]
	storagenodestr := r.Form["storagenodelist"][0]
	metanodestr := r.Form["storagenodelist"][0]
	managenodelist := strings.Split(managenodestr,",")
	storagenodelist := strings.Split(storagenodestr,",")
	metanodelist := strings.Split(metanodestr,",")
	cluster := &Cluster{
		ClusterIP:clusterip,
		StorageNodeList:storagenodelist,
		ManageNodeList:managenodelist,
		MetaNodeList:metanodelist,
	}

	err := cluster.InitCluster()
	if err != nil{

	}
	err = cluster.CreateManageNode()
	if err != nil{

	}

	err = cluster.CreateMetaNode()
	if err != nil{

	}

	err = cluster.CreateClientNode()
	if err != nil{
		response := types.ErrorResponse{err.Error(),-1}
		httputils.WriteJSON(w, http.StatusOK,&response)
	}
	return httputils.WriteJSON(w, http.StatusOK,"test post interface")
}

func (t *clusterRouter)addStorageNode(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK,"test post interface")
}

func (t *clusterRouter)addMetaNode(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK,"test post interface")
}

func (t *clusterRouter)addClientNode(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK,"test post interface")
}

func (t *clusterRouter)getStorageStatus(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK,"test post interface")
}

func (t *clusterRouter)getMetaStatus(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK,"test post interface")
}

func (t *clusterRouter)getClientStatus(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK,"test post interface")
}