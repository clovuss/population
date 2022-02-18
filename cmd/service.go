package main

import (
	"fmt"
	"github.com/clovuss/population/preparedata"
)

func (app *application) Updatedatabase() {
	in, out := preparedata.CompareFiles()
	for _, prikrep := range in {
		err := app.Repo.InsertOne(prikrep, "main")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	for _, prikrep := range out {
		err := app.Repo.InsertOne(prikrep, "out")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = app.Repo.DeleteByOne(prikrep.ENP)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}
