package messages

import (
	"encoding/json"
	"net/http"
	"noox/utils"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /sendmessage", h.sendmessage)
}

func (h *Handler) sendmessage(w http.ResponseWriter, r *http.Request) {
	utils.GetSession(r)
	var m utils.MesagePayload
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return
	}
	response := map[string]interface{}{
		"success": "true",
		"Message": m,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
