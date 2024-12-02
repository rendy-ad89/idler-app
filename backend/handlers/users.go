package handlers

import (
	"context"
	"fmt"
	db "idler/app/sqlc"
	"idler/app/util"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type SaveProgressRequest struct {
	ID   int64
	Cash float64
}

func CreateUser(c *gin.Context) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	defer conn.Close(ctx)
	queries := db.New(conn)

	var req db.CreateUserParams
	if util.ValidateUserRequest(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := queries.GetUserByUsername(ctx, req.Username)
	if err != nil && err != pgx.ErrNoRows {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	if user.Username == req.Username {
		c.JSON(http.StatusBadRequest, "Username "+user.Username+" already taken")
		return
	}

	newUser, err := queries.CreateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	machines, err := queries.GetMachines(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	var batchReq []db.CreateUsersMachinesParams
	for _, e := range machines {
		createReq := db.CreateUsersMachinesParams{
			UserID:    newUser.ID,
			MachineID: e.ID,
		}
		batchReq = append(batchReq, createReq)
	}

	results := queries.CreateUsersMachines(ctx, batchReq)
	results.Exec(func(i int, err error) {
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error")
			return
		}
	})

	c.JSON(http.StatusOK, "Registered successfully! Please proceed to login.")
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

func SaveProgress(c *gin.Context) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	defer conn.Close(ctx)
	queries := db.New(conn)

	var req SaveProgressRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	var cash pgtype.Numeric
	if err := cash.Scan(fmt.Sprintf("%.2f", req.Cash)); err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	v, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	claims := v.(*util.Claims)
	saveReq := db.SaveProgressParams{
		ID:   claims.ID,
		Cash: cash,
	}
	if err := queries.SaveProgress(ctx, saveReq); err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func CalcOfflineProfits(c *gin.Context) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	defer conn.Close(ctx)
	queries := db.New(conn)

	v, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	claims := v.(*util.Claims)

	user, err := queries.GetUser(ctx, claims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	usersAmplifiers, err := queries.GetUsersAmplifiers(ctx, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	usersGenerators, err := queries.GetUsersGenerators(ctx, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	var offlineTime = math.Floor(time.Since(user.LastSave.Time).Seconds())
	var amplification = 100.0
	var offlineCash = 0.0
	for _, e := range usersAmplifiers {
		generation, err := e.Generation.Float64Value()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error")
			return
		}
		amplification += generation.Float64 * float64(e.Level.Int32)
	}
	for _, e := range usersGenerators {
		generation, err := e.Generation.Float64Value()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error")
			return
		}
		offlineCash += generation.Float64 * float64(e.Level.Int32) * offlineTime * amplification / 100
	}

	cash, err := user.Cash.Float64Value()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	totalCash := cash.Float64 + offlineCash
	var updatedCash pgtype.Numeric
	if err := updatedCash.Scan(fmt.Sprintf("%.2f", totalCash)); err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	user.Cash = updatedCash

	req := db.SaveProgressParams{
		ID:   user.ID,
		Cash: user.Cash,
	}
	if err := queries.SaveProgress(ctx, req); err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}

	c.JSON(http.StatusOK, updatedCash)
}
