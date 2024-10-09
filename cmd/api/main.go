package main

import (
	"log"
	config "vibex-api/configs"
	"vibex-api/internal/routes"
)

func main() {
	//load environment
	config.LoadEnv()

	//postgress connection
	config.ConnectDB()

	router := routes.SetupRouter()

	log.Fatal(router.Run(":8080"))
}
