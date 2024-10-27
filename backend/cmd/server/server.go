package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

var connections *pgx.Conn

func main() {
	router := gin.Default()

	connectionConfig := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "hive",
		Password: "password",
		User:     "admin",
	}
	connection, err := pgx.Connect(connectionConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	router.POST("/users", func(context *gin.Context) {
		type Body struct {
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}

		var userDetails Body

		if err := context.BindJSON(&userDetails); err != nil {
			fmt.Println("Failed to bind JSON:", err.Error()) // Print the error message
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		fmt.Printf("User: %+v\n", userDetails)
		// createdTime := time.Now()

		hashedPassword, err := argon2id.CreateHash("pa$$word", argon2id.DefaultParams)

		if err != nil {
			log.Fatal(err)
		}

		sqlStatement := `
		INSERT INTO users (email, password_hash, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

		commandTag, err := connection.Exec(sqlStatement, userDetails.Email, hashedPassword, userDetails.Firstname, userDetails.Lastname)
		if err != nil {
			fmt.Println("Failed to execute insert query:", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		fmt.Println("Rows affected ", commandTag.RowsAffected())

		context.JSON(http.StatusOK, gin.H{
			"message": "User data received",
			"body":    userDetails,
		})

	})

	router.GET("/users", func(context *gin.Context) {

		context.JSON(http.StatusOK, gin.H{
			"message": "hit",
		})
	})

	router.POST("/user", func(context *gin.Context) {
		name := context.Param("name")
		fmt.Println(name)
		// c.JSON(http.StatusOk, gin.H{})
	})

	router.POST("/permission/request", function(context *gin.Context) {
		
	})

	router.Run()
}
