package main

import (
	"go-vue-app/backend/internal/database"
	"go-vue-app/backend/pkg/routers"
	"go-vue-app/backend/pkg/seeds"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := database.GetMysqlConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	if os.Getenv("ENABLED_MIGRATIONS") == "1" {
		database.AutoMigrate(db)
		seeds.Up(db)
	}

	app := fiber.New()
	routers.RegisterRoutes(db, app)
	app.Listen(":3000")
}
