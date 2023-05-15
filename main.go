package main

import (
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname,omitempty"`
	Mail     string `json:"mail"`
}

var database map[int]User

const (
	listenAddr = ":2200"
	httpPrefix = "/api"
)

func init() {
	database = make(map[int]User)
	database[0] = User{
		Name:     "Thim Cederlund",
		Nickname: "thimc",
		Mail:     "xxxxxxxxxxxxxx",
	}
}

func main() {
	log.SetPrefix("[INFO] ")
	http.HandleFunc(httpPrefix+"/create", handleCreate)
	http.HandleFunc(httpPrefix+"/update", handleUpdate)
	http.HandleFunc(httpPrefix+"/delete", handleDelete)
	http.HandleFunc(httpPrefix+"/read", handleRead)

	log.Printf("Serving on %s\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
