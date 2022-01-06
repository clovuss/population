package view

import (
	"fmt"
	"github.com/clovuss/population/models"
	"html/template"
	"net/http"
	"reflect"
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
	Adress       string
	PrikAuto     bool
	PrikDate     bool
	SnilsDoc     string
	UchZav       bool
	Pacients     []*models.Pacient
	PacientsView [][]interface{}
	//formatter func(time time.Time) string
}

func Construct(params map[string][]string, uch int, pac []*models.Pacient) *View {
	if len(pac) == 0 {
		return nil
	}
	pacsview := make([][]interface{}, 0)
	for _, v := range pac {
		vl := reflect.ValueOf(*v)
		if vl.Kind() != reflect.Struct {
			return nil
		}
		out := make([]interface{}, vl.NumField())
		for i := 0; i < vl.NumField(); i++ {
			out[i] = vl.Field(i).Interface()
		}
		pacsview = append(pacsview, out)
	}

	v := &View{
		NumUch:       uch,
		Fio:          "on",
		Gender:       true,
		Pacients:     pac,
		PacientsView: pacsview,
	}
	for k, _ := range params {
		switch k {
		case "adress":
			v.Adress = "on"
		case "enp":
			v.Enp = true
		case "birthday":
			v.Birthday = true
		case "snils":
			v.Snils = true
		case "prikreptype":
			v.PrikAuto = true
		case "prikrepdate":
			v.PrikDate = true

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

func (v *View) RenderHTML(w http.ResponseWriter) {
	tem, err := template.ParseFiles(
		"./ui/html/base.gohtml",
		"./ui/html/ranger.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	err = tem.Execute(w, v)
	if err != nil {
		fmt.Println(err)
	}

}
