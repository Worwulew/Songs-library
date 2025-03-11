package main

import (
	"context"
	"fmt"
	"github.com/Worwulew/Songs-library/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("Usage: %s COMMAND\n", os.Args[0]))
	}

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
	command := os.Args[1]
	migrationDir, err := filepath.Abs("migrations")
	if err != nil {
		panic(fmt.Sprintf("Cannot resolve absolute path: %s", err))
	}
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlDbname,
		cfg.Postgres.PostgresqlPassword,
	)

	db, err := goose.OpenDBWithDriver(cfg.Postgres.PgDriver, dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("goose: failed to open DB: %v", err))
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = goose.RunContext(ctx, command, db, migrationDir); err != nil {
		panic(fmt.Sprintf("goose %s: %v", command, err))
	}

	log.Println("Migration command executed successfully")
}
