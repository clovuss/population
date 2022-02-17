package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/uch/", app.viewbyUch)
	router.HandleFunc("/enp/", app.viewbyEnp)
	router.HandleFunc("/editenp/", app.editbyEnp)
	fileserver := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileserver))

	return router
}
