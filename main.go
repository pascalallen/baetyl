package main

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/pascalallen/Baetyl/src/Adapter/Http/JSend"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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
		v1.GET("/test", handleTest)
		v1.POST("/upload", handleFileUpload)

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

func handleRegisterUser(c *gin.Context) {}

func handleLoginUser(c *gin.Context) {}

func handleLogoutUser(c *gin.Context) {}

func handleRefreshUserSession(c *gin.Context) {}

func handleRequestPasswordReset(c *gin.Context) {}

func handleResetPassword(c *gin.Context) {}

func handleFileUpload(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.IndentedJSON(
			http.StatusMethodNotAllowed,
			JSend.FailResponse[string]{
				Status: "fail",
				Data:   "Method now allowed",
			},
		)

		return
	}

	// Restrict uploaded file size to be 20MB or less
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			JSend.ErrorResponse[string]{
				Status:  "error",
				Message: err.Error(),
			},
		)
	}

	file, m, err := c.Request.FormFile("file")
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			JSend.ErrorResponse[string]{
				Status:  "error",
				Message: err.Error(),
			},
		)
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				JSend.FailResponse[string]{
					Status: "fail",
					Data:   err.Error(),
				},
			)
		}
	}(file)

	// TODO: Process uploaded file

	successMessage := fmt.Sprintf("Successfully processed file: %s", m.Filename)
	c.IndentedJSON(
		http.StatusCreated,
		JSend.SuccessResponse[string]{
			Status: "success",
			Data:   successMessage,
		},
	)
}

func handleTest(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		JSend.SuccessResponse[any]{
			Status: "success",
			Data:   nil,
		},
	)
}
