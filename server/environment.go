package server

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func (a API) LoadEnvVariables() error {
	var filename string
	env := os.Getenv("BASEAPIENV")

	// Default file is the local one
	if env == "" {
		filename = ".env.local"
		gin.SetMode(gin.DebugMode)
	} else if env == "prod" {
		filename = ".env"
		gin.SetMode(gin.ReleaseMode)
	} else {
		filename = ".env."+env
		gin.SetMode(gin.TestMode)
	}

	err := godotenv.Overload(filename)
	return err
}
