package utils

type RegisterPayload struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type LoginPayload struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
