package main

import (
	"log"

	"github.com/joho/godotenv"

	"myutilityx.com/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
		return
	}
	routes := routes.RegisterRoutes()
	err = routes.Run(":8080")
	if err != nil {
		log.Fatal("failed to start the server!!")
		return
	}
}
