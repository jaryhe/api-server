package test

import (
	"golang.org/x/net/context"
	"net/http"
	"fmt"
	"storm/api/server/httputils"
	"github.com/sirupsen/logrus"
)

type TestOut struct {
	name string `json:",omitempty"`
	age  string `json:",omitempty"`
}
func (t *testRouter)testGet(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	fmt.Println("testget")

	m := make(map[string]int)
	m["ddddddddddddddd"] = 1
	m["aaaaaaaaaaaaaaa"] = 2
	name := vars["name"]
	fmt.Println(name)
	if err := httputils.ParseForm(r); err != nil {
		return err
	}
	fmt.Println(r.Form)
	defer r.Body.Close()
	buf := make([]byte, 1024)
	for {
		_, err := r.Body.Read(buf)
		if err != nil {
			break
		}
	}
	logrus.Infoln(name)
	fmt.Println(buf)
	testOut :=&TestOut{
		name:"longyu.he",
		age:"30",
	}
	return httputils.WriteJSON(w, http.StatusOK,testOut)
}

func (t *testRouter)testPost(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	fmt.Println("testpost")
	buf := make([]byte, 1024)
	for {
		_, err := r.Body.Read(buf)
		if err != nil {
			break
		}
	}
	fmt.Println(buf)
	return httputils.WriteJSON(w, http.StatusCreated,"test post interface")
}

func (t *testRouter)testDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	fmt.Println("testdelete")

	return httputils.WriteJSON(w, http.StatusProcessing,"test delete interface")
}