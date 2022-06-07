package main

import (
	"github.com/jabutech/simple-blog/router"
	"github.com/jabutech/simple-blog/util"
)

func main() {

	// Open connection to db
	db := util.SetupDb()
	// Router
	server := router.NewRouter(db)

	// Run router
	server.Run()
}
