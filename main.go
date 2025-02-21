package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/lambaharsh01/surveyItBackend/database"
	"github.com/lambaharsh01/surveyItBackend/middleware"
	"github.com/lambaharsh01/surveyItBackend/routes"
	"github.com/lambaharsh01/surveyItBackend/utils"
)

func main() {

	loadEnvError := godotenv.Load(".env")
	if loadEnvError != nil {
		log.Fatal("Error accessing .env file")
	}

	port := utils.GetEnv("PORT")

	dbInstance := database.InitDb()

	router := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))

	router.Use(middleware.RateLimitHandler())
	router.Use(middleware.ErrorHandler())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Server Listening",
		})

		var CurrentTimeStamp string
		if err := dbInstance.Raw("SELECT CURRENT_TIMESTAMP FROM DUAL").Scan(&CurrentTimeStamp).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": true,
				"message": "Database Not Listening At All",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": fmt.Sprintf("Database Listening at: %s", CurrentTimeStamp),
		})

	})

	apiGroup := router.Group("/api")

	authGroup := apiGroup.Group("/auth")
	routes.AuthRoutes(authGroup, dbInstance)

	meGroup := apiGroup.Group("/user")
	meGroup.Use(middleware.AuthenticationHandler())
	routes.MeRoutes(meGroup, dbInstance)

	fetchGroup := apiGroup.Group("/fetch")
	fetchGroup.Use(middleware.AuthenticationHandler())
	routes.FetchRoutes(fetchGroup, dbInstance)

	router.Run(":" + port)

}
