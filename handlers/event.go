package handlers

import (
	"context"
	"log"
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
	db, ok := c.Get("db").(*pgx.Conn)
	if !ok || db == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection not found"})
	}

	event := new(Event)
	if err := c.Bind(event); err != nil {
		log.Printf("Error binding event: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	log.Printf("Received event: %+v", event)

	// Ensure data integrity
	if event.Title == "" || event.Date == "" || event.UserID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}

	_, err := db.Exec(context.Background(),
		"INSERT INTO events (title, description, date, user_id) VALUES ($1, $2, $3, $4)",
		event.Title, event.Description, event.Date, event.UserID)
	if err != nil {
		log.Printf("Error inserting event into database: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "Event created successfully"})
}

func GetEvents(c echo.Context) error {
	log.Println("GetEvents")
	db := c.Get("db").(*pgx.Conn)
	rows, err := db.Query(context.Background(), "SELECT * FROM events")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var events []Event
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		events = append(events, event)
	}

	return c.JSON(http.StatusOK, events)
}
