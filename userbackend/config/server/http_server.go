package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	serverdto "test/userbackend/config/server/dto"
)

type HttpServer struct {
	port string
}

func Construct(port string) *HttpServer {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return &HttpServer{port: port}
}

func (server *HttpServer) StartListening() {
	var apiMux *http.ServeMux = http.NewServeMux()

	fmt.Println("testing testing")
	fmt.Println("TEST END")

	apiMux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		server.handler(w, r)
	})

	if err := http.ListenAndServe(server.port, apiMux); err != nil {
		fmt.Println("HTTP server error:", err)
	}
}

func (server *HttpServer) handler(w http.ResponseWriter, r *http.Request) {
	var uri string = strings.TrimPrefix(r.URL.Path, "/api")
	fmt.Println("test ==================== ")
	fmt.Println(uri)

	var callback *serverdto.Route = resolveRoute(uri)
	var controllMethod func() serverdto.Response = callback.ControllerMethod
	var result serverdto.Response = controllMethod()

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonBytes)
}
