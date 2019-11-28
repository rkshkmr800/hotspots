package main

import (
	"github.com/rkshkmr800/hotspots/backend/main-framework/server"
)

func main() {

	router := server.CreateRouter()
	server.StartServer(router)

}
