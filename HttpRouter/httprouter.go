package main

import (
	"time"

	"LearnGo/HttpRouter/server"
)

func main() {
	go server.StartServer()
	time.Sleep(time.Minute * 3)
}
