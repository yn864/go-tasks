package handlers

import (
	"math/rand"
	"net/http"
	"time"
)

func HardOp(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)+10))
	w.WriteHeader(http.StatusOK)
}
