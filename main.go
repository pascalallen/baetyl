package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dunglas/mercure"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pascalallen/Baetyl/src/Adapter/Database"
	RegisterUserAction "github.com/pascalallen/Baetyl/src/Adapter/Http/Action/Api/V1/Auth"
	GormPermissionRepository "github.com/pascalallen/Baetyl/src/Adapter/Repository/Auth/Permission"
	GormRoleRepository "github.com/pascalallen/Baetyl/src/Adapter/Repository/Auth/Role"
	GormUserRepository "github.com/pascalallen/Baetyl/src/Adapter/Repository/Auth/User"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
	"log"
	"net/http"
	"os"
)

func main() {
	unitOfWork, err := Database.NewGormUnitOfWork()
	if err != nil {
		log.Fatal(err)
	}

	if err := unitOfWork.AutoMigrate(&Permission.Permission{}, &Role.Role{}, &User.User{}); err != nil {
		err := fmt.Errorf("failed to auto migrate database: %s", err)
		log.Fatal(err)
	}

	var permissionRepository Permission.PermissionRepository = GormPermissionRepository.GormPermissionRepository{
		UnitOfWork: unitOfWork,
	}
	var roleRepository Role.RoleRepository = GormRoleRepository.GormRoleRepository{
		UnitOfWork: unitOfWork,
	}
	var userRepository User.UserRepository = GormUserRepository.GormUserRepository{
		UnitOfWork: unitOfWork,
	}
	dataSeeder := Database.DataSeeder{
		UnitOfWork:           unitOfWork,
		PermissionRepository: permissionRepository,
		RoleRepository:       roleRepository,
		UserRepository:       userRepository,
	}
	if err := dataSeeder.Seed(); err != nil {
		log.Fatal(err)
	}

	mercureHub, err := mercure.NewHub(
		mercure.WithPublisherJWT([]byte(os.Getenv("MERCURE_JWT_KEY")), "HS256"),
		mercure.WithSubscriberJWT([]byte(os.Getenv("MERCURE_JWT_KEY")), "HS256"),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer mercureHub.Stop()
	http.Handle("/.well-known/mercure", mercureHub)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/public/assets", "./public/assets")

	environment := map[string]string{
		"APP_BASE_URL":       os.Getenv("APP_BASE_URL"),
		"MERCURE_PUBLIC_URL": os.Getenv("MERCURE_PUBLIC_URL"),
	}
	envJson, _ := json.Marshal(environment)
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"environment": base64.StdEncoding.EncodeToString(envJson),
		})
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
