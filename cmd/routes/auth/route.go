package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"noox/db"
	"noox/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	coll := db.Client.Database("users").Collection("users")
	var storedUser utils.RegisterPayload
	err = coll.FindOne(context.TODO(), bson.M{"username": u.Username}).Decode(&storedUser)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	//Compare password
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	token, err := utils.GenerateJWT()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
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
	//hash the password
	hashp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = coll.InsertOne(context.TODO(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, `{"message":"User created successfully"}`)
}
