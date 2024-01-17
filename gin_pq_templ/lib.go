package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type App struct {
	db *pgxpool.Pool
}

func NewApp() *App {
	return &App{db: setupDatabase()}
}

func (app *App) ApiSampleSetup(g *gin.Engine) {
	setupPrometheus(g)
	g.GET("/api/samples", GetSamples)
}

func setupPrometheus(r *gin.Engine) {
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
}

func GetSamples(c *gin.Context) {

}

func setupDatabase() *pgxpool.Pool {
	databaseUrl := os.Getenv("DATABASE_URL")
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		panic(err)
	}
	return dbPool
}
