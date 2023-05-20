package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeResponse(res http.ResponseWriter, httpCode int, response ApiResponse) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)

	if err := json.NewEncoder(res).Encode(&response); err != nil {
		log.Println(err)
	}
}
