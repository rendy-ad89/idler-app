package handlers

import (
	"context"
	db "idler/app/sqlc"
	"idler/app/util"
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

func GetUsersMachines(c *gin.Context) {
	conn, err := pgx.Connect(c, os.Getenv("DB_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	defer conn.Close(c)
	queries := db.New(conn)

	v, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	claims := v.(*util.Claims)

	usersMachines, err := queries.GetUsersMachines(c, claims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	c.JSON(http.StatusOK, usersMachines)
}

func UpdateUsersMachines(c *gin.Context) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	defer conn.Close(ctx)
	queries := db.New(conn)

	var req []db.UpdateUsersMachinesParams
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	v, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	claims := v.(*util.Claims)
	if req[0].UserID != claims.ID {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	results := queries.UpdateUsersMachines(ctx, req)
	results.Exec(func(i int, err error) {
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error")
			return
		}
	})

	c.JSON(http.StatusOK, "OK")
}
