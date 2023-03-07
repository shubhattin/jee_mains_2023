package main

import (
	"api/jee"
	process_data "api/jee/process_data"
	"api/kry"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		process_data.GetData("", "")
		return
	}

	// loading .env file
	godotenv.Load()

	if kry.PROD_ENV {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := gin.Default()
	set_cors_headers(app)

	app.Use(gzip.Gzip(gzip.DefaultCompression))

	jee.GetRoutes(app)

	PORT, IS_PORT := os.LookupEnv("PORT")
	if kry.DEV_ENV && !IS_PORT {
		app.Run(":3428")
		// use nodemon to restart the server on file change
		// nodemon --exec go run main.go --signal SIGTERM
	} else if IS_PORT {
		app.Run(":" + PORT)
	} else {
		app.Run(":8080")
	}
}

func set_cors_headers(router *gin.Engine) {
	MAX_AGE := 24 * time.Hour

	origins := []string{"http://localhost"}
	STATIC_SITE_URL, IS_STATIC_SITE_URL := os.LookupEnv("STATIC_SITE_URL")
	// Allow CORS for static site
	if IS_STATIC_SITE_URL {
		origins = append(origins, STATIC_SITE_URL)
	}

	config := cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           MAX_AGE,
	})
	router.Use(config)
}
