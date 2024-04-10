package tools

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Err string `json:"error,omitempty"`
}

type Sucsess struct {
	Id uint64 `json:"banner_id,omitempty"`
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

func SendSucsessContent(w http.ResponseWriter, code int, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	//json.NewEncoder(w).Encode(body)
	w.Write([]byte(body))
}

func SendSucsessArray(w http.ResponseWriter, code int, arr interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(arr)
	//w.Write([]byte(body))
}

func SendSucsessId(w http.ResponseWriter, code int, id uint64) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	//json.NewEncoder(w).Encode(body)
	//w.Write([]byte(body))
	json.NewEncoder(w).Encode(&Sucsess{Id: id})
}
