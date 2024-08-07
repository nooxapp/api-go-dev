package token

import (
	"net/http"
	"noox/utils"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /token", h.token)
}

func (h *Handler) token(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Write([]byte("You are authenticated" + "\n"))
}
