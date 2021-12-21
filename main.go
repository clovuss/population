package main

import (
	"context"
	"fmt"
	"github.com/clovuss/population/models"
	"github.com/clovuss/population/preparedata"
	"github.com/jackc/pgx/v4/pgxpool"
)

type application struct {
	popXML map[string]preparedata.PRIKREP
	LinkDB *models.PacientDB
}

func main() {
	var dsn = "postgres://user:0405@localhost:5432/population"
	dbpool := openDb(dsn)
	defer dbpool.Close()
	app := application{
		LinkDB: &models.PacientDB{DB: dbpool}}
	app.popXML = preparedata.ImprooveDataPrepare("14")
	w, err := app.LinkDB.GetByPid("65")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println((*w).Birthday)

	//map15 := improoveDataPrepare("15")
	//
	//for k, _ := range map15 {
	//
	//	_, ok := map14[k]
	//	if !ok {
	//		fmt.Println(k)
	//	}
	//}

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
