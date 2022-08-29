package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pascalallen/Baetyl/src/Adapter/Database"
	RegisterUserAction "github.com/pascalallen/Baetyl/src/Adapter/Http/Action/Api/V1/Auth"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
	"log"
	"net/http"
)

func init() {
	unitOfWork := Database.GormUnitOfWork{}
	if err := unitOfWork.InitDbSession(); err != nil {
		panic(err.Error())
	}

	db := unitOfWork.DatabaseSession
	if err := db.AutoMigrate(&Permission.Permission{}, &Role.Role{}, &User.User{}); err != nil {
		panic("failed to migrate database")
	}

	// temp for debugging
	dataSeeder := Database.DataSeeder{}
	if err := dataSeeder.Seed(); err != nil {
		panic(err.Error())
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
