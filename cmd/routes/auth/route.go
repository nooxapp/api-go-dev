package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"noox/db"
	"noox/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/login", h.Login)
	mux.HandleFunc("POST /auth/register", h.Register)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var u utils.LoginPayload
	w.Header().Set("content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//find user
	coll := db.Client.Database("users").Collection("users")
	err = coll.FindOne(context.TODO(), u).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var u utils.RegisterPayload
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//id shit so I don't have OBJECK IDS ex: _id ObjectID("607598484768416672")
	objectID := primitive.NewObjectID()
	u.ID = objectID.Hex()
	coll := db.Client.Database("users").Collection("users")
	_, err = coll.InsertOne(context.TODO(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, `{"message":"User created successfully"}`)
}
