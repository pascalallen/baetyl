package main

import (
	"fmt"
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
