package main

import (
	"fmt"
	"log"

	"github.com/A-mey/GO-AUTH/cmd/routes"
	"github.com/A-mey/GO-AUTH/config"
	"github.com/A-mey/GO-AUTH/libs/sql"
	"github.com/gin-gonic/gin"
)

func main() {

	config := config.LoadConfig()
	env := config.AppEnv

	r := gin.Default()
	sql.InitializeSQL()

	routes.InitializeRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		fmt.Println("Running in", env, "Mode")
	}
}
