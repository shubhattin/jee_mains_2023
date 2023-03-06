package main

import (
	"api/jee"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/gzip"
)

func main() {
	// loading .env file
	godotenv.Load()
	app := gin.Default()
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	jee.GetRoutes(app)

	// app.GET("/api/result_count", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	PORT, IS_PORT := os.LookupEnv("PORT")
	if !IS_PORT {
		app.Run(":3428")
		// use nodemon to restart the server on file change
		// nodemon --exec go run main.go --signal SIGTERM
	} else {
		app.Run(":" + PORT)
	}
}
