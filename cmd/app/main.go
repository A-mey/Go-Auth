package main

import (
	"fmt"
	"log"

	userRoutes "github.com/A-mey/Auth_db/api/health/routes"
	healthRoutes "github.com/A-mey/Auth_db/api/v1/users/routes"
	"github.com/A-mey/Auth_db/config/config"
	"github.com/A-mey/Auth_db/config/sql"
	"github.com/gin-gonic/gin"
)

func main() {

	config := config.LoadConfig()
	env := config.AppEnv

	r := gin.Default()
	sql.InitializeSQL()

	userRoutes.InitializeRoutes(r)
	healthRoutes.InitializeRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		fmt.Println("Running in", env, "Mode")
	}
}
