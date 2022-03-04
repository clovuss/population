package main

import (
	"context"
	"fmt"
	"github.com/clovuss/population/models"
	"github.com/clovuss/population/preparedata"
	"github.com/clovuss/population/view"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
	"time"
)

type application struct {
	popXML       map[string]preparedata.PRIKREP
	Repo         *models.PacientDB
	snilsdoc     map[int][]string
	View         *view.View
	LastUpdateDB time.Time
	Quantity     int
}

func main() {

	var dsn = "postgres://user01:@192.168.88.3:5432/population"
	dbpool := openDb(dsn)
	defer dbpool.Close()
	app := application{
		Repo: &models.PacientDB{DB: dbpool},
		snilsdoc: map[int][]string{
			1:  {"037-431-051 26"},
			2:  {"173-024-614 35"},
			3:  {"037-431-166 36"},
			4:  {"139-926-189 10"},
			5:  {"121-876-460 62"},
			6:  {"037-431-155 33"},
			7:  {"037-431-143 29"},
			8:  {"037-431-139 33"},
			9:  {"036-579-107 76"},
			10: {"037-431-141 27"},
			11: {"171-395-174 75"},
			12: {"037-431-161 31"},
		},
	}

	var err error

	app.LastUpdateDB, err = app.Repo.GetLAstUpdate()
	if err != nil {
		fmt.Println(err)
	}
	app.Quantity, err = app.Repo.GetLQuantity()

	if err != nil {
		fmt.Println(err)
	}
	server := http.Server{
		Addr:      ":8080",
		Handler:   app.routes(),
		TLSConfig: nil,
		ErrorLog:  nil,
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}

func openDb(dsn string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Println(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return pool
}
