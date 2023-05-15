package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeResponse(res http.ResponseWriter, httpCode int, response ApiResponse) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	_, err = res.Write(jsonData)
	if err != nil {
		log.Println(err)
		return
	}
}
