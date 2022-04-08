package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./views/chart.html")
	if err != nil {
		log.Println(err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func StartWebServer() {
	http.HandleFunc("/chart", viewChartHandler)
	http.ListenAndServe(":8000", nil)
}
