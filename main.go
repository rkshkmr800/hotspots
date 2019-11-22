package main

import . "github.com/rkshkmr800/hotspots/server"

func main() {

	router := server.CreateRouter(db)
	server.StartServer(router)

}
