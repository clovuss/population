package main

import (
	"fmt"
	"github.com/clovuss/population/view"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, req *http.Request) {
	//tem, err := template.ParseFiles("./ui/html/base.html")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//r, err := app.Repo.GetByPid("65")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = tem.Execute(w, r)
	//if err != nil {
	//	fmt.Println(err)
	//}
	tem, err := template.ParseFiles("./ui/html/ex.html")

	if err != nil {
		fmt.Println(err)
	}
	sl := make(map[int][]string, 5)
	sl[1] = []string{"раз1", "раз2", "раз3", "раз4"}
	sl[2] = []string{"два1", "два2"}
	sl[3] = []string{}
	sl[4] = []string{"четыре"}
	sl[5] = []string{"пять", "пять", "пять", "пять", "пять"}

	err = tem.Execute(w, sl)
	if err != nil {
		fmt.Println(err)
	}

}
func (app *application) view(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	paramsUser := make(map[string][]string)
	paramsUser["uch_zav"] = []string{"on"}

	numberUch, err := strconv.Atoi(req.URL.Path[len("/view/"):])
	if err != nil {
		fmt.Println(err)
	}
	if numberUch < 1 || numberUch > 12 {
		http.NotFound(w, req)
		return
	}
	//fmt.Println("numberuch: ", app.snilsdoc[numberUch][1])
	pac, err := app.Repo.GetByUch(paramsUser, app.snilsdoc[numberUch])

	if err != nil {
		fmt.Println("mis", err)
	}

	//fmt.Println("exit:", *(pac[2]))

	if err != nil {
		fmt.Println(err)
	}

	vs := view.Construct(paramsUser, numberUch, pac)

	vs.RenderHTML(w)
	if err != nil {
		fmt.Println(err)
	}

}
