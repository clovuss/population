package main

import (
	"fmt"
	"github.com/clovuss/population/view"
	"net/http"
	"strconv"
)

//home returns  start page
func (app *application) home(w http.ResponseWriter, req *http.Request) {
	//tem, err := template.ParseFiles(
	//	"./ui/html/base.gohtml",
	//	"./ui/html/ranger.gohtml")
	_, err := w.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	}

	//err = tem.Execute(w, "Hello from home")
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func (app *application) viewbyUch(w http.ResponseWriter, req *http.Request) {
	//get view options from user
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	paramsUser := make(map[string][]string)
	paramsUser = req.Form
	fmt.Println(paramsUser)

	numberUch, err := strconv.Atoi(req.URL.Path[len("/uch/"):])
	if err != nil {
		fmt.Println(err)
	}
	if numberUch < 1 || numberUch > 12 {
		http.NotFound(w, req)
		return
	}

	resDb, err := app.Repo.GetByUch(paramsUser, app.snilsdoc[numberUch])

	if err != nil {
		fmt.Println("mistake from handler", err)
	}

	if err != nil {
		fmt.Println(err)
	}

	vs := view.Construct(paramsUser, numberUch, resDb)

	vs.RenderHTMLUch(w)
	if err != nil {
		fmt.Println(err)
	}

}
func (app *application) Vr(enp string) (*view.View, error) {
	p, err := app.Repo.GetByEnp(enp)
	if err != nil {
		fmt.Println(err)
	}
	v := &view.View{Pacient: *p}
	return v, nil
}

func (app *application) viewbyEnp(w http.ResponseWriter, req *http.Request) {

	enpfromuser := req.URL.Path[len("/enp/"):]
	if len(enpfromuser) != 16 {
		http.NotFound(w, req)
		return
	}
	enp, err := strconv.Atoi(enpfromuser)
	rz := strconv.Itoa(enp)
	if err != nil {
		fmt.Println(err)
	}
	app.View, err = app.Vr(rz)

	if err != nil {
		return
	}
	app.View.RenderHTMLEnp(w)

}

func (app *application) editbyEnp(w http.ResponseWriter, req *http.Request) {
	enpfromuser := req.URL.Path[len("/editenp/"):]
	if len(enpfromuser) != 16 {
		http.NotFound(w, req)
		return
	}
	enp, err := strconv.Atoi(enpfromuser)
	//обработать ошибку

	if err != nil {
		return
	}

	rz := strconv.Itoa(enp)

	if app.View == nil {
		app.View, err = app.Vr(rz)
	}

	//fmt.Println(app.View)
	app.View.RenderHTMLEditEnp(w)

}
