package view

import (
	"fmt"
	"github.com/clovuss/population/models"
	"html/template"
	"net/http"
	"time"
)

type View struct {
	NumUch       int
	Pid          string
	Fio          string
	Birthday     bool
	Gender       bool
	Enp          bool
	Snils        bool
	Placeofbirth string
	RegionName   string
	Adress       bool
	LiveAdress   bool
	Document     bool
	PrikAuto     bool
	PrikDate     bool
	SnilsDoc     string
	UchZav       bool
	CardNum      bool
	Phone        bool
	Pacients     []*models.Pacient
	Pacient      models.Pacient
}

func Construct(params map[string][]string, uch int, pac []*models.Pacient) *View {
	v := &View{
		NumUch:   uch,
		Fio:      "on",
		Enp:      true,
		Pacients: pac,
		//PacientsView: pacsview,
	}
	for k, _ := range params {
		switch k {
		case "adress":
			v.Adress = true
		case "live_adress":
			v.LiveAdress = true
		case "gender":
			v.Gender = true
		case "birthday":
			v.Birthday = true
		case "snils":
			v.Snils = true
		case "prikreptype":
			v.PrikAuto = true
		case "prikrepdate":
			v.PrikDate = true
		case "phone":
			v.Phone = true
		case "document":
			v.Document = true
		case "card_num":
			v.CardNum = true
		case "uch_zav":
			v.UchZav = true
		}
	}
	return v
}

func (v View) DateFormat(t time.Time) string {
	return t.Format("02-01-2006")

}
func (v View) GenderView(g string) string {
	if g == "1" {
		return "М"
	}
	return "Ж"
}
func (v View) PrikrepView(p string) string {
	if p == "1" {
		return "Т"
	}
	return "З"
}

func (v View) Adder(i int) int {
	return i + 1
}

func (v *View) RenderHTMLUch(w http.ResponseWriter) {
	//tem, err := template.ParseFiles(
	//	"./html/generator.gohtml",
	//	"./html/main.gohtml")
	//tem, err := template.ParseFiles(
	//	"./html/1.gohtml",
	//
	//	"./html/2.gohtml",
	//
	//	"./html/3.gohtml")

	tem, err := template.ParseFiles(
		"./html/startranger.gohtml",

		"./html/main.gohtml",

		"./html/generator.gohtml",
		"./html/end.gohtml")

	if err != nil {
		fmt.Println(err)
	}
	err = tem.Execute(w, v)
	if err != nil {
		fmt.Println(err)
	}

}
func (v *View) RenderHTMLEnp(w http.ResponseWriter) {
	tem, err := template.ParseFiles(
		"./html/startenp.gohtml",
		"./html/main.gohtml",
		"./html/byenp.gohtml",
		"./html/end.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	err = tem.Execute(w, v)
	if err != nil {
		fmt.Println(err)
	}

}
func (v *View) RenderHTMLEditEnp(w http.ResponseWriter) {
	tem, err := template.ParseFiles(
		"./html/starteditenp.gohtml",
		"./html/main.gohtml",
		"./html/editbyenp.gohtml",
		"./html/end.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	err = tem.Execute(w, v)
	if err != nil {
		fmt.Println(err)
	}

}
