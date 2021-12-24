package main

import (
	"fmt"
	"github.com/clovuss/population/view"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, req *http.Request) {
	tem, err := template.ParseFiles("./ui/html/base.html")
	if err != nil {
		fmt.Println(err)
	}
	r, err := app.LinkDB.GetByPid("65")
	if err != nil {
		fmt.Println(err)
	}

	err = tem.Execute(w, r)
	if err != nil {
		fmt.Println(err)
	}

}
func (app *application) view(w http.ResponseWriter, req *http.Request) {
	//tem, err := template.ParseFiles("./ui/html/base.html",
	//	"./ui/html/ranger.html")
	//if err != nil {
	//	fmt.Println(err)
	//}
	req.ParseForm()
	numberUch, err := strconv.Atoi(req.URL.Path[len("/view/"):])
	if err != nil {
		fmt.Println(err)
	}
	if numberUch < 1 || numberUch > 12 {
		http.NotFound(w, req)
		return
	}
	rec, err := app.LinkDB.GetByUch(app.snilsdoc[numberUch])
	v := view.View{
		Record: rec,
		NumUch: numberUch,
		Fio:    "on",
	}
	v.RenderHTML(w, &v)

	if err != nil {
		fmt.Println(err)
	}
	//err = tem.Execute(w, p)
	//if err != nil {
	//	fmt.Println(err)
	//}

}
