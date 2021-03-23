package controllers

import (
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("views/index.html"))
	res := &Page{
		Title: os.Getenv("APP_NAME"),
	}
	tpl.Execute(w, res)
}
