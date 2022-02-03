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

func (app *application) view(w http.ResponseWriter, req *http.Request) {
	//get view options from user
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	paramsUser := make(map[string][]string)
	paramsUser = req.Form
	fmt.Println(paramsUser)

	numberUch, err := strconv.Atoi(req.URL.Path[len("/view/"):])
	if err != nil {
		fmt.Println(err)
	}
	if numberUch < 1 || numberUch > 12 {
		http.NotFound(w, req)
		return
	}
	fmt.Println("numberUch:", numberUch)
	//pac, err := app.Repo.GetByUch(paramsUser, app.snilsdoc[numberUch])

	if err != nil {
		fmt.Println("mis", err)
	}

	//fmt.Println("exit:", *(pac[2]))

	if err != nil {
		fmt.Println(err)
	}

	vs := view.Construct(paramsUser, numberUch, nil)
	fmt.Println(vs.NumUch)

	vs.RenderHTML(w)
	if err != nil {
		fmt.Println(err)
	}

}
