package main

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/pascalallen/Baetyl/src/Adapter/Http/JSend"
	"github.com/pascalallen/Baetyl/src/Domain/PasswordHash"
	"github.com/pascalallen/Baetyl/src/Domain/User"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RegisterUserFormValidations struct {
	FirstName       string `form:"first_name" json:"first_name" binding:"required"`
	LastName        string `form:"last_name" json:"last_name" binding:"required"`
	EmailAddress    string `form:"email_address" json:"email_address" binding:"required,email"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
}

func main() {
	connStr := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(db)

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
			auth.POST("/register", handleRegisterUser)
			auth.POST("/session", handleLoginUser)
			auth.DELETE("/session", handleLogoutUser)
			auth.PATCH("/session", handleRefreshUserSession)
			auth.POST("/reset", handleRequestPasswordReset)
			auth.POST("/password", handleResetPassword)
		}
	}

	log.Fatal(router.Run(":80"))
}

func handleRegisterUser(c *gin.Context) {
	var form RegisterUserFormValidations

	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			JSend.FailResponse[string]{
				Status: "fail",
				Data:   err.Error(),
			},
		)

		return
	}

	user := User.Register(form.FirstName, form.LastName, form.EmailAddress)
	passwordHash := PasswordHash.Create(form.Password)
	user.SetPasswordHash(passwordHash)

	c.JSON(
		http.StatusCreated,
		JSend.SuccessResponse[User.User]{
			Status: "success",
			Data:   *user,
		},
	)
}

func handleLoginUser(c *gin.Context) {}

func handleLogoutUser(c *gin.Context) {}

func handleRefreshUserSession(c *gin.Context) {}

func handleRequestPasswordReset(c *gin.Context) {}

func handleResetPassword(c *gin.Context) {}
