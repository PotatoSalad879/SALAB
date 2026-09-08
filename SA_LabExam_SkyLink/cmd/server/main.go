package main

import (
	"log"

	"github.com/joho/godotenv"

	"skylink/internal/db"
)

// SkyLink — schema bootstrap entrypoint.
// รันเพื่อสร้าง/อัปเดต schema จริงลง PostgreSQL (ต้อง `docker compose up -d` ก่อน)
func main() {
	_ = godotenv.Load()

	gdb, err := db.Open(db.DSN())
	if err != nil {
		log.Fatalf("connect: %v", err)
	}

	if err := db.AutoMigrate(gdb); err != nil {
		log.Fatalf("auto-migrate: %v", err)
	}

	log.Println("SkyLink schema migrated successfully")
}
