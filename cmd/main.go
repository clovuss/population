package main

import (
	"context"
	"fmt"
	"github.com/clovuss/population/models"
	"github.com/clovuss/population/preparedata"
	"github.com/clovuss/population/view"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type application struct {
	popXML   map[string]preparedata.PRIKREP
	Repo     *models.PacientDB
	snilsdoc map[int][]string
	View     *view.View
}

func main() {
	var dsn = "postgres://user:0405@localhost:5432/population"
	dbpool := openDb(dsn)
	defer dbpool.Close()
	app := application{
		Repo: &models.PacientDB{DB: dbpool},
		snilsdoc: map[int][]string{
			1:  {"037-431-051 26"},
			2:  {"173-024-614 35", "171-395-174 75"},
			3:  {"037-431-166 36"},
			4:  {"139-926-189 10"},
			5:  {"121-876-460 62"},
			6:  {"037-431-155 33"},
			7:  {"037-431-143 29"},
			8:  {"037-431-139 33"},
			9:  {"036-579-107 76"},
			10: {"037-431-141 27"},
			11: {"171-395-174 75", "173-024-614 35"},
			12: {"037-431-161 31"},
		},
	}

	//mapafromXML := preparedata.ImprooveDataPrepareDate("11")
	//fmt.Println(len(mapafromXML))

	//mapafromuch := preparedata.Prepare1csv("1")
	//for _, v := range mapafromuch {
	//	err := app.Repo.InsertUch(v)
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//}

	//fmt.Println("длина в м")

	//os.Exit(0)
	server := http.Server{
		Addr:      ":8080",
		Handler:   app.routes(),
		TLSConfig: nil,
		ErrorLog:  nil,
	}

	err := server.ListenAndServe()
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
