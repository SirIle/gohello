package main // import "github.com/sirile/gohello"

import (
	"net/http"
	"os"
	"time"

	"github.com/emicklei/go-restful"
)

type Message struct {
	Hostname string
	Time     time.Time
	Language string
}

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(hello))
	ws.Route(ws.HEAD("/").To(header))
	restful.Add(ws)
	http.ListenAndServe(":80", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
	name, err := os.Hostname()
	if err == nil {
		msg := Message{name, time.Now(), "go"}
		resp.WriteAsJson(msg)
	}
}

func header(req *restful.Request, resp *restful.Response) {
	resp.WriteHeader(200)
}
