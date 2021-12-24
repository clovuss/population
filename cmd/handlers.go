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
	r, err := app.Repo.GetByPid("65")
	if err != nil {
		fmt.Println(err)
	}

	err = tem.Execute(w, r)
	if err != nil {
		fmt.Println(err)
	}

}
func (app *application) view(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	paramsUser := make(map[string]string, 0)
	paramsUser["checkBirthday"] = req.Form.Get("birthday")
	paramsUser["checkEnp"] = req.Form.Get("enp")
	paramsUser["checkSnils"] = req.Form.Get("snils")
	paramsUser["checkAdress"] = req.Form.Get("adress")
	paramsUser["checkPrikreptype"] = req.Form.Get("prikreptype")
	paramsUser["checkPrikrepdate"] = req.Form.Get("prikrepdate")
	numberUch, err := strconv.Atoi(req.URL.Path[len("/view/"):])
	if err != nil {
		fmt.Println(err)
	}
	if numberUch < 1 || numberUch > 12 {
		http.NotFound(w, req)
		return
	}
	vs := view.Construct(paramsUser, numberUch)
	vs.RenderHTML(w)
	if err != nil {
		fmt.Println(err)
	}

}
