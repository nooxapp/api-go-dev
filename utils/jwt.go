package utils

import (
	"net/http"
	"time"
)

func GenerateJWT(r *http.Request) {
	extime := time.Now().Add(time.Minute * 15)
}
