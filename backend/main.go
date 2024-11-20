package main

import (
	"log"
	"nabbr-appraisal/api"
	"nabbr-appraisal/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize and check for command line flags
	utils.InitCommandLineFlags()

	// Load environment variables
	err := godotenv.Load(utils.EnvFilePath)
	if err != nil {
		log.Fatalf("error - Error loading .env file: %s", err)
	}

	// Connect to the databases
	if os.Getenv("DB_SQLITE_ENABLE") == "true" {
		utils.LoadDbConnectToSqlite()
	}

	// Start webserver
	api.StartServer()
}