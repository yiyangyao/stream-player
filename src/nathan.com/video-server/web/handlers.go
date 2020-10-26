package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := &HomePage{Name: "nathan"}
	t, e := template.ParseFiles("./templates/home.html")
	if e != nil {
		log.Printf("parsing template home.html error: %s", e)
		return
	}

	t.Execute(w, p)
}
