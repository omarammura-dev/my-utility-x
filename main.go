package main

import (
	"myutilityx.com/routes"
)

func main() {

	routes := routes.RegisterRoutes()
	routes.Run(":8080")
}
