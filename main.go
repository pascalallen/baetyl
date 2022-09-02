package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	RegisterUserAction "github.com/pascalallen/Baetyl/src/Adapter/Http/Action/Api/V1/Auth"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/public/assets", "./public/assets")
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", RegisterUserAction.Handle)
			auth.POST("/session", handleLoginUser)
			auth.DELETE("/session", handleLogoutUser)
			auth.PATCH("/session", handleRefreshUserSession)
			auth.POST("/reset", handleRequestPasswordReset)
			auth.POST("/password", handleResetPassword)
		}
	}

	log.Fatal(router.Run(":80"))
}

func handleLoginUser(c *gin.Context) {}

func handleLogoutUser(c *gin.Context) {}

func handleRefreshUserSession(c *gin.Context) {}

func handleRequestPasswordReset(c *gin.Context) {}

func handleResetPassword(c *gin.Context) {}
