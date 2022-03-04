package main

import (
	"fmt"
	"github.com/clovuss/population/preparedata"
	"github.com/clovuss/population/view"
	"net/http"
	"strconv"
	"time"
)

//home returns  start page
func (app *application) home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		http.Redirect(w, req, "/uch/1", 303)
	}
}

func (app *application) viewbyUch(w http.ResponseWriter, req *http.Request) {
	//get view options from user
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	paramsUser := make(map[string][]string)
	paramsUser = req.Form

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

	quantityUch := len(resDb)
	vs := app.View.Construct(paramsUser, numberUch, resDb, quantityUch, app.Quantity, app.LastUpdateDB)
	if err != nil {
		fmt.Println(err)
	}
	vs.RenderHTMLUch(w)
	if err != nil {
		fmt.Println(err)
	}

}

func (app *application) viewPacient(enp string) (*view.View, error) {
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
	if req.Method == http.MethodPost {
		err = req.ParseForm()
		if err != nil {
			return
		}
		rawNumberPhone := req.PostForm.Get("phone")
		tel, err := preparedata.NewtelValidator(rawNumberPhone)
		if err != nil {
			fmt.Println(err)
		}
		t := strconv.Itoa(tel)
		err = app.Repo.UpdatePhone(t, enpfromuser)
		if err != nil {
			fmt.Println(err)
		}

	}

	app.View, err = app.viewPacient(rz)

	if err != nil {
		return
	}
	app.View.RenderHTMLEnp(w)

}

func (app *application) update(w http.ResponseWriter, req *http.Request) {
	//поискали файлы.  Возващается дата формирования самого свежего файла
	files, err := preparedata.SearchFiles()
	if err != nil {
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}
	d := time.Hour * 24
	dateOfLastFile := files[0].ModTime().AddDate(0, 0, -1)

	if len(files) == 1 {
		if dateOfLastFile.Sub(app.LastUpdateDB) < d {
			//sincelastupdate := time.Since(dateOfLastFile)

			//_, err := w.Write([]byte("сведения из базы данных соответствуют загруженным файлам. Обновление было"+fmt.Fprintf(sincelastupdate.Truncate(24 * time.Hour).Hours())))
			fmt.Fprint(w, "сведения из базы данных соответствуют загруженным файлам. Обновление сделано с файла, скачанного ", time.Since(files[0].ModTime()).Hours(), " часов назад.")
			if err != nil {
				return
			}
			return
		} else {
			_, err := w.Write([]byte("большое расхождение между датой xml-файла и датой обновления бд"))
			if err != nil {
				return
			}
			return
		}
	}
	err = app.Updatedatabase()
	if err != nil {
		fmt.Fprint(w, err)

		return
	}
	app.LastUpdateDB, err = app.Repo.GetLAstUpdate()
	if err != nil {
		fmt.Fprint(w, err)
	}
	app.Quantity, err = app.Repo.GetLQuantity()
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, "сведения обновлены на ", app.LastUpdateDB, ". Население: ", app.Quantity, " детей.")

}

func (app *application) find(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}
	paramsUser := make(map[string][]string)
	paramsUser = req.Form
	isepmty := true
	for _, strings := range paramsUser {
		if strings[0] != "" {
			isepmty = false
		}
	}

	fmt.Println(paramsUser)
	if isepmty {
		fmt.Fprint(w, "введите данные для поиска")
		return
	}
	res, err := app.Repo.FindByName(paramsUser)

	v := &view.View{Pacients: res}

	v.RenderHTMLUch(w)

	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(len(res))
}
