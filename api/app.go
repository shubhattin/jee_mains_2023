package main

import (
	"api/jee"
	process_data "api/jee/process_data"
	"api/kry"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		data, _ := process_data.GetData("", "")
		process_data.GetMetaData("")
		process_data.GetResult(&data)
		return
	}

	if kry.DEV_ENV {
		// loading .env file
		godotenv.Load()
	}

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		AppName:       "JEE Mains 2023 Score Calculator",
	})

	set_cors_headers(app)

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	jee.GetRoutes(app)

	PORT, IS_PORT := os.LookupEnv("PORT")
	if kry.DEV_ENV && !IS_PORT {
		app.Listen(":3428")
	} else if IS_PORT {
		app.Listen(":" + PORT)
	} else {
		app.Listen(":8080")
	}
}

func set_cors_headers(app *fiber.App) {
	MAX_AGE := 24 * time.Hour

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("STATIC_SITE_URL"),
		AllowHeaders: "*",
		MaxAge:       int(MAX_AGE.Seconds()),
	}))
}
