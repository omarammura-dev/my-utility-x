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
	routes.Run(":8080")
}
