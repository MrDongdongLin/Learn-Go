package main

import (
	"log"
	"net/http"

	"LearnGo/redis/server"
)

func main() {
	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8001", router))
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
