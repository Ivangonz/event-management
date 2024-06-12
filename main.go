package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ivangonz/event-management/handlers"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	e.GET("/api/events", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "List of events",
		})
	})

	e.POST("/api/register", handlers.RegisterUser)
	e.POST("/api/events", handlers.CreateEvent)
	e.GET("/api/events", handlers.GetEvents)

	e.Logger.Fatal(e.Start(":3000"))
}
