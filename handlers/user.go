package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c echo.Context) error {
	db := c.Get("db").(*pgx.Conn)
	user := new(User)

	if err := c.Bind(user); err != nil {
		log.Printf("Error binding user: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	_, err := db.Exec(context.Background(),
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		user.Name, user.Email, user.Password)
	if err != nil {
		log.Printf("Error inserting user into database: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, user)
}
