package main

import (
	"RESTApi/go-rest-api/config"
)

func main() {
	config.Init()
	config.InitDB()
	config.HandleRequests()
}
