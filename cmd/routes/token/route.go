package token

import (
	"context"
	"encoding/json"
	"net/http"
	"noox/db"
	"noox/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /token", h.token)
}

func (h *Handler) token(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userID := claims.ID
	coll := db.Client.Database("users").Collection("users")
	filter := bson.D{{Key: "_id", Value: userID}}
	projection := bson.D{
		{Key: "username", Value: 1},
	}

	var result bson.M
	err = coll.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to convert result to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("You are authenticated " + userID))
	w.Write(resultJSON)
}
