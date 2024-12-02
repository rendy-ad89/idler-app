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

type AuthResponse struct {
	AccessToken string  `json:"accessToken"`
	ID          int64   `json:"id"`
	Username    string  `json:"username"`
	Cash        float64 `json:"cash"`
}

func ValidateAuth(c *gin.Context) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	defer conn.Close(ctx)
	queries := db.New(conn)

	var reqUser db.User
	if err := c.BindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := queries.GetUserByUsername(ctx, reqUser.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid credentials.")
		return
	}

	if user.Password != reqUser.Password {
		c.JSON(http.StatusUnauthorized, "Invalid credentials.")
		return
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	cash, err := user.Cash.Float64Value()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	data := AuthResponse{
		AccessToken: token,
		ID:          user.ID,
		Username:    user.Username,
		Cash:        cash.Float64,
	}
	c.JSON(http.StatusOK, data)
}
