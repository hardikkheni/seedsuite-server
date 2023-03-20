package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hardikkheni/seedsuite-server/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	app := gin.Default()
	routes.LoadWebRoutes(app.Group("/"))
	routes.LoadApiRoutes(app.Group("/api"))
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}
	log.Fatal(app.Run(fmt.Sprintf(":%d", port)))
}
