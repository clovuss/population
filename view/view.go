package view

import (
	"fmt"
	"github.com/clovuss/population/models"
	"html/template"
	"net/http"
)

type View struct {
	Record       []*models.Pacient
	NumUch       int
	Pid          string
	Enp          string
	Fio          string
	Birthday     string
	Gender       string
	Snils        string
	Placeofbirth string
	RegionName   string
	City         string
	NasPunkt     string
	Street       string
	House        string
	Korp         string
	Kvart        string
	PrikAuto     string
	PrikDate     string
	SnilsDoc     string
}

func (v *View) RenderHTML(w http.ResponseWriter, s *View) {
	tem, err := template.ParseFiles("./ui/html/base.html",
		"./ui/html/ranger.html")
	if err != nil {
		fmt.Println(err)
	}

	tem.Execute(w, s)

}
