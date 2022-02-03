package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/view/", app.view)
	fileserver := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileserver))

	return router
}
