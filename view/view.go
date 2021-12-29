package view

import (
	"fmt"
	"github.com/clovuss/population/models"
	"html/template"
	"net/http"
)

type View struct {
	Pacients     []*models.Pacient
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

func Construct(params map[string][]string, uch int, pac []*models.Pacient) *View {
	v := &View{
		NumUch:   uch,
		Fio:      "on",
		Gender:   "on",
		Pacients: pac,
	}
	for k, _ := range params {
		switch k {
		case "adress":
			v.Adress = "on"
		case "enp":
			v.Enp = "on"
		case "birthday":
			v.Birthday = "on"
		case "snils":
			v.Snils = "on"
		case "prikreptype":
			v.PrikAuto = "on"
		case "prikrepdate":
			v.PrikDate = "on"

		}
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
