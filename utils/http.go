package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

//
// ErrorResponder will respond with an error
// to the given http.ResponseWriter
//
func ErrorResponder(w http.ResponseWriter, statusCode int, givenError error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{
		Status: "error",
		Data:   map[string]string{"error": givenError.Error()},
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

//
// JSONResponder will respond with the given
// json to the given http.ResponseWriter
//
func JSONResponder(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status: "success",
		Data:   data,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

//
// JSONSuccessResponder will respond with the given
// json to the given http.ResponseWriter
//
func JSONSuccessResponder(w http.ResponseWriter) {
	var data interface{}
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status: "success",
		Data:   data,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

//
// CorsHandler provides cors support
//
func CorsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Add("Access-Control-Allow-Headers", "*")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Write(nil)
		} else {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Add("Access-Control-Allow-Headers", "*")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			h.ServeHTTP(w, r)
		}
	}
}
