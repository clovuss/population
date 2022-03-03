package main

import (
	"fmt"
	"github.com/clovuss/population/preparedata"
	"os"
	"time"
)

func (app *application) Updatedatabase() error {
	in, out, err := preparedata.Dbservice()
	if err != nil {
		return err
	}
	f, err := preparedata.SearchFiles()
	if err != nil {
		return err
	}
	var outdate time.Time
	outdate = f[0].ModTime()
	if outdate.Before(app.LastUpdateDB) {
		err := fmt.Errorf("сведения в бд актуальнее файлов, проверьте их")
		return err
	}

	for _, prikrep := range in {
		err := app.Repo.InsertIntoMain(prikrep)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	//таблица убывших. Дата убытия = дата файла
	for _, prikrep := range out {
		err := app.Repo.InsertIntoOut(prikrep, outdate)

		if err != nil {
			fmt.Println("добавляем в аут", err, prikrep)
			return err
		}
		err = app.Repo.DeleteByOne(prikrep.ENP)
		if err != nil {

			fmt.Println(err)
			return err
		}

	}

	err = os.Remove("./data/" + f[1].Name())
	if err != nil {
		return err
	}
	return nil
}
