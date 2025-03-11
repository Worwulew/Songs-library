package main

import (
	"github.com/Worwulew/Songs-library/config"
	"github.com/Worwulew/Songs-library/internal/server"
	"github.com/Worwulew/Songs-library/pkg/db/postgres"
	"github.com/Worwulew/Songs-library/pkg/logger"
	"github.com/joho/godotenv"
	"log"
)

// @title Songs library REST API
// @version 1.0
// @description Songs library REST API
// @contact.name Grigory
// @contact.url https://github.com/Worwulew
// @contact.email gregoridagmo@gmail.com
// @BasePath /api/v1
func main() {
	log.Println("Starting api server")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	srv := server.NewServer(cfg, psqlDB, appLogger)
	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}

	if err := appLogger.SugarLogger.Sync(); err != nil {
		appLogger.SugarLogger.Error(err)
	}
}
