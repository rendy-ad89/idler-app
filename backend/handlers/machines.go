package handlers

import (
	"context"
	db "idler/app/sqlc"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetMachines(c *gin.Context) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	defer conn.Close(ctx)
	queries := db.New(conn)

	machines, err := queries.GetMachines(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	c.JSON(http.StatusOK, machines)
}
