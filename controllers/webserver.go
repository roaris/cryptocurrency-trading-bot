package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/cryptocurrency-trading-bot/models"
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

func candleHandler(w http.ResponseWriter, r *http.Request) {
	candles, err := models.GetAllCandles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(candles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func StartWebServer() {
	http.HandleFunc("/chart", viewChartHandler)
	http.HandleFunc("/api/candles", candleHandler)
	http.ListenAndServe(":8000", nil)
}
