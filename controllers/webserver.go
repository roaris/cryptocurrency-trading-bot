package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

	df := models.DataFrameCandle{Candles: candles}

	sma := r.URL.Query().Get("sma")
	if sma != "" {
		strSmaPeriod1 := r.URL.Query().Get("smaPeriod1")
		strSmaPeriod2 := r.URL.Query().Get("smaPeriod2")
		strSmaPeriod3 := r.URL.Query().Get("smaPeriod3")
		smaPeriod1, err := strconv.Atoi(strSmaPeriod1)
		if err == nil {
			df.AddSma(smaPeriod1)
		}
		smaPeriod2, err := strconv.Atoi(strSmaPeriod2)
		if err == nil {
			df.AddSma(smaPeriod2)
		}
		smaPeriod3, err := strconv.Atoi(strSmaPeriod3)
		if err == nil {
			df.AddSma(smaPeriod3)
		}
	}

	res, err := json.Marshal(df)
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
