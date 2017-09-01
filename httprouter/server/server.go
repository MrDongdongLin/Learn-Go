package server

import (
	"log"
	"net/http"

	"LearnGo/HttpRouter/controller"
	"github.com/julienschmidt/httprouter"
)

func StartServer() {
	router := httprouter.New()
	router.GET("/get/:requestid", controller.GetTaskProcess)
	router.POST("/post/:requestid", controller.PostTaskProcess)
	log.Fatal(http.ListenAndServe(":8002", router))
}
