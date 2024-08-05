package api

import (
	"fmt"
	"net/http"
	"noox/cmd/routes/test"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	messageService := test.NewHandler()
	messageService.RegisterRoutes(router)

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	fmt.Println("Listening on http://localhost" + s.addr + "/api/v1/")
	return http.ListenAndServe(s.addr, router)
}
