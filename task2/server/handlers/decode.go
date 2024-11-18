package handlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

type InputString struct {
	Input string `json:"inputString"`
}

type OutputString struct {
	Output string `json:"outputString"`
}

func Decode(w http.ResponseWriter, r *http.Request) {
	var clientMessage InputString
	err := json.NewDecoder(r.Body).Decode(&clientMessage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decodedRawBytes, err := base64.StdEncoding.DecodeString(clientMessage.Input)
	if err != nil {
		log.Fatal()
	}

	serverResponse := OutputString{
		Output: string(decodedRawBytes),
	}

	err = json.NewEncoder(w).Encode(serverResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
