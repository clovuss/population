package view

import (
	"fmt"
	"html/template"
	"net/http"
)

type View struct {
	NumUch       int
	Pid          string
	Enp          string
	Fio          string
	Birthday     string
	Gender       string
	Snils        string
	Placeofbirth string
	RegionName   string
	Adress       string
	PrikAuto     string
	PrikDate     string
	SnilsDoc     string
}

func Construct(params map[string]string, uch int) *View {
	v := &View{
		NumUch:   uch,
		Fio:      "on",
		Birthday: params["checkBirthday"],
		Enp:      params["checkEnp"],
		Snils:    params["checkSnils"],
		Adress:   params["checkAdress"],
		PrikDate: params["checkPrikrepdate"],
		PrikAuto: params["checkPrikreptype"],
	}
	return v
}
func (v *View) RenderHTML(w http.ResponseWriter) {
	tem, err := template.ParseFiles(
		"./ui/html/base.html",
		"./ui/html/ranger.html")
	if err != nil {
		fmt.Println(err)
	}

	err = tem.Execute(w, v)
	if err != nil {
		fmt.Println(err)
	}

}
