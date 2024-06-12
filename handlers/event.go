package handlers

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
)

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	UserID      int    `json:"user_id"`
}

func CreateEvent(c echo.Context) error {
	db := c.Get("db").(*pgx.Conn)
	event := new(Event)

	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err := db.Exec(context.Background(),
		"INSERT INTO events (title, description, date, user_id) VALUES ($1, $2, $3, $4)",
		event.Title, event.Description, event.Date, event.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}
