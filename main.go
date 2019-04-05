package main

import (
	"os"

	service "github.com/horis233/GoGo-Service/service"
	"github.com/cloudfoundry-community/go-cfenv"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, _ := cfenv.Current()
	server := service.NewServer(appEnv)
	server.Run(":" + port)
}