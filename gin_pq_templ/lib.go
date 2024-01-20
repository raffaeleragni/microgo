package main

import (
	"context"
	"log"
	"net/http"
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
	g.GET("/api/samples", app.GetSamples)
}

func setupPrometheus(r *gin.Engine) {
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
}

type Sample struct {
	Id   string
	Name string
}

func (app *App) GetSamples(c *gin.Context) {
	rows, err := app.db.Query(context.Background(), "select * from sample")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []Sample
	for rows.Next() {
		var r Sample
		err := rows.Scan(&r.Id, &r.Name)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
}

func setupDatabase() *pgxpool.Pool {
	databaseUrl := os.Getenv("DATABASE_URL")
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		panic(err)
	}
	return dbPool
}
