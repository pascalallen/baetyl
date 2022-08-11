package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	RegisterUserAction "github.com/pascalallen/Baetyl/src/Adapter/Http/Action/Api/V1/Auth"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&Permission.Permission{}, &Role.Role{}, &User.User{}); err != nil {
		panic("failed to migrate database")
	}
}

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
