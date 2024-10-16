package handlers

import (
	"fmt"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "v1.0.0")
}
