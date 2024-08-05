package auth

import (
	"fmt"
	"net/http"
	"noox/utils"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/login", h.Login)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var u utils.LoginPayload
	w.Header().Set("Content-Type", "application/json")
	err := utils.ReadJSON(r, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(u)
}
