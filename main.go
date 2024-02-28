package main

import (
	"myutilityx.com/db"
	"myutilityx.com/routes"
)

func main() {
	err := db.Init("links")
	if err != nil {
		panic(err)
	}
	routes := routes.RegisterRoutes()
	routes.Run(":8080")
}
