package management

import (
	"fmt"
	"os"
	"io"
	"os/exec"
	"io/ioutil"
)

var drbdFile = "/etc/drbd.d/mdata.res"

type  Cluster struct{
	ClusterIP   string
	NodeList    []string
	HostIP      map[string]string
}

func (c *Cluster)InitCluster() error{
	return nil
}
func (c *Cluster)SetHost() error{
	// set /etc/hosts file
	return nil
}

func (c *Cluster)IpGetName() error{
	return nil
}

func (c *Cluster)CheckClsterIP() error{
	// check cluster ip is in use
	return nil
}

func (c *Cluster)AuthCluster() error{
	// cluster authorization
	// pcs cluster auth host1 host2

	return nil
}

func (c *Cluster)SetupCluster() error{
	// setup cluster
	// pcs cluster setup --name clustername host1 host2
	return nil
}

func (c *Cluster) CreateDrdbConfFile() error{
	drdbConfStr := fmt.Sprintf(drbdTemplate,"192.168.10.180","192.168.10.181")
	f,err1 := os.OpenFile(drbdFile,os.O_CREATE|os.O_WRONLY,0666)
	if err1 != nil{
		return err1
	}
	_,err2 := io.WriteString(f, drdbConfStr)
    if err2 != nil{
    	return err2
	}
	return nil
}

func (c *Cluster) CreateDrdbDevice() error{
	cmd := exec.Command("drbdadm","create-md","mdata")
	stdout,err :=cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start();err != nil {
		return err
	}
	stderr ,_ := cmd.StderrPipe()

	errBytes,_ := ioutil.ReadAll(stderr)
	opBytes,err := ioutil.ReadAll(stdout)
	fmt.Println(string(opBytes))
	fmt.Println(string(errBytes))
	return nil
}

func (c *Cluster)CreateClusterResource()error{
	return nil
}

func (c *Cluster)CreateConstraintColocation()error{
	return nil
}

func (c *Cluster)CreateConstraintOrder()error{
	return nil
}

func (c *Cluster)CheckClusterStatus() error{
	return nil
}

func (c *Cluster)GetRsourceStatus() error{
	return nil
}



var drbdTemplate = `resource mdata {
protocol C;
meta-disk internal;
device /dev/drbd1;
syncer {
verify-alg sha1;
}
net {
allow-two-primaries;
}
on %s {
disk /dev/centos/mlv;
address %s:7789;
}
on %s {
disk /dev/centos/mlv;
address %s:7789;
}
}`