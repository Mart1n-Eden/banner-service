package tools

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Err string `json:"error,omitempty"`
}

func SendStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func SendError(w http.ResponseWriter, code int, err string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&Error{Err: err})
	//fmt.Println(w)
}

func SendSucsess(w http.ResponseWriter, code int, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	//json.NewEncoder(w).Encode(body)
	w.Write([]byte(body))
}
