package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dunglas/mercure"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pascalallen/baetyl/src/Adapter/Database"
	RegisterUserAction "github.com/pascalallen/baetyl/src/Adapter/Http/Action/Api/V1/Auth"
	GormPermissionRepository "github.com/pascalallen/baetyl/src/Adapter/Repository/Auth/Permission"
	GormRoleRepository "github.com/pascalallen/baetyl/src/Adapter/Repository/Auth/Role"
	GormUserRepository "github.com/pascalallen/baetyl/src/Adapter/Repository/Auth/User"
	"github.com/pascalallen/baetyl/src/Domain/Auth/Permission"
	"github.com/pascalallen/baetyl/src/Domain/Auth/Role"
	"github.com/pascalallen/baetyl/src/Domain/Auth/SecurityToken"
	"github.com/pascalallen/baetyl/src/Domain/Auth/User"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	unitOfWork, err := Database.NewGormUnitOfWork()
	if err != nil {
		log.Fatal(err)
	}

	migrate(unitOfWork)

	var permissionRepository Permission.PermissionRepository = GormPermissionRepository.GormPermissionRepository{
		UnitOfWork: unitOfWork,
	}
	var roleRepository Role.RoleRepository = GormRoleRepository.GormRoleRepository{
		UnitOfWork: unitOfWork,
	}
	var userRepository User.UserRepository = GormUserRepository.GormUserRepository{
		UnitOfWork: unitOfWork,
	}

	seed(unitOfWork, permissionRepository, roleRepository, userRepository)

	router := gin.Default()
	mercureHub := createMercureHub()

	registerRoutes(router, mercureHub)

	log.Fatal(router.Run(":80"))
}

func migrate(unitOfWork *gorm.DB) {
	if err := unitOfWork.AutoMigrate(&Permission.Permission{}, &Role.Role{}, &User.User{}, &SecurityToken.SecurityToken{}); err != nil {
		err := fmt.Errorf("failed to auto migrate database: %s", err)
		log.Fatal(err)
	}
}

func seed(unitOfWork *gorm.DB, permissionRepository Permission.PermissionRepository, roleRepository Role.RoleRepository, userRepository User.UserRepository) {
	dataSeeder := Database.DataSeeder{
		UnitOfWork:           unitOfWork,
		PermissionRepository: permissionRepository,
		RoleRepository:       roleRepository,
		UserRepository:       userRepository,
	}
	if err := dataSeeder.Seed(); err != nil {
		log.Fatal(err)
	}
}

func createMercureHub() *mercure.Hub {
	mercureHub, err := mercure.NewHub(
		mercure.WithPublisherJWT([]byte(os.Getenv("MERCURE_JWT_KEY")), "HS256"),
		mercure.WithSubscriberJWT([]byte(os.Getenv("MERCURE_JWT_KEY")), "HS256"),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer func(mercureHub *mercure.Hub) {
		if err := mercureHub.Stop(); err != nil {
			log.Fatal(err)
		}
	}(mercureHub)

	return mercureHub
}

func registerRoutes(router *gin.Engine, mercureHub *mercure.Hub) {
	router.LoadHTMLGlob("templates/*")
	router.Static("/public/assets", "./public/assets")
	// TODO: Determine how to publish updates to Mercure hub
	router.Any("/.well-known/mercure", func(context *gin.Context) {
		// Mercure publish data: id, topic, data
		mercureHub.PublishHandler(context.Writer, context.Request)
		mercureHub.SubscribeHandler(context.Writer, context.Request)
	})

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
}

// TODO: Implement session registerRoutes
func handleLoginUser(c *gin.Context) {
	log.Print(c)
}

func handleLogoutUser(c *gin.Context) {
	log.Print(c)
}

func handleRefreshUserSession(c *gin.Context) {
	log.Print(c)
}

func handleRequestPasswordReset(c *gin.Context) {
	log.Print(c)
}

func handleResetPassword(c *gin.Context) {
	log.Print(c)
}
