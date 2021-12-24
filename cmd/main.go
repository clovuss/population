package main

import (
	"context"
	"fmt"
	"github.com/clovuss/population/models"
	"github.com/clovuss/population/preparedata"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type application struct {
	popXML   map[string]preparedata.PRIKREP
	LinkDB   *models.PacientDB
	snilsdoc map[int][]string
}

func main() {
	var dsn = "postgres://user:0405@localhost:5432/population"
	dbpool := openDb(dsn)
	defer dbpool.Close()
	app := application{
		LinkDB: &models.PacientDB{DB: dbpool},
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
	app.popXML = preparedata.ImprooveDataPrepare("14")
	server := http.Server{
		Addr:      ":8080",
		Handler:   app.routes(),
		TLSConfig: nil,
		ErrorLog:  nil,
	}

	//map15 := improoveDataPrepare("15")
	//
	//for k, _ := range map15 {
	//
	//	_, ok := map14[k]
	//	if !ok {
	//		fmt.Println(k)
	//	}
	//}
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