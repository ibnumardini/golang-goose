package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

const (
	dbPath         = "./db/demo.db"
	migrationsPath = "./migrations"
	driver         = "sqlite"
)

func main() {
	if err := os.MkdirAll("./db", 0755); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(driver, dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect(driver); err != nil {
		log.Fatalf("set dialect: %v", err)
	}

	command := "up"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "up":
		if err := goose.Up(db, migrationsPath, goose.WithAllowMissing()); err != nil {
			log.Fatalf("goose up: %v", err)
		}
	case "down":
		if err := goose.Down(db, migrationsPath); err != nil {
			log.Fatalf("goose down: %v", err)
		}
	case "reset":
		if err := goose.Reset(db, migrationsPath); err != nil {
			log.Fatalf("goose reset: %v", err)
		}
	case "status":
		if err := goose.Status(db, migrationsPath); err != nil {
			log.Fatalf("goose status: %v", err)
		}
	case "version":
		version, err := goose.GetDBVersion(db)
		if err != nil {
			log.Fatalf("goose version: %v", err)
		}
		fmt.Printf("current version: %d\n", version)
	default:
		fmt.Printf("unknown command: %s\nusage: go run main.go [up|down|reset|status|version]\n", command)
		os.Exit(1)
	}
}
