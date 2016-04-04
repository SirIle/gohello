package main // import "github.com/sirile/gohello"

import (
	"net/http"
	"os"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("go-image-test")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Message defines the structure of the returned JSON
type Message struct {
	Hostname string    `json:"hostname"`
	Time     time.Time `json:"time"`
	Language string    `json:"language"`
}

func main() {
	logging.SetBackend(logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(hello))
	ws.Route(ws.HEAD("/").To(header))
	restful.Add(ws)
	log.Debug("Starting the server and listening for requests")
	http.ListenAndServe(":80", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
	name, err := os.Hostname()
	if err == nil {
		msg := Message{name, time.Now(), "go"}
		resp.PrettyPrint(false)
		resp.WriteAsJson(msg)
	}
	log.Debug("JSON served")
}

// Implemented so that HAProxy heartbeat checks work properly
func header(req *restful.Request, resp *restful.Response) {
	resp.WriteHeader(200)
}
