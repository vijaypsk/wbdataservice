package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func ShowData(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, ctries)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	//fmt.Println(ctries)
	if err := json.NewEncoder(w).Encode(indicators); err != nil {
		panic(err)
	}
}

func GetIndicatorsForYear(w http.ResponseWriter, r *http.Request) {
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	var year int64
	var err error
	if year, err = strconv.ParseInt(vars["year"], 10, 64); err != nil {
		panic(err)
	}
	results := GetIndicatorsFor(year)
	fmt.Println(results)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
	return

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func GetIndicatorDataForYear(w http.ResponseWriter, r *http.Request) {
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	var year int64
	var icode string
	var err error
	if year, err = strconv.ParseInt(vars["year"], 10, 64); err != nil {
		panic(err)
	}
	icode = vars["icode"]
	results := GetIndicatorDataFor(year, icode)
	fmt.Println(results)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
	return

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func GetIndicatorDefnitions(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, ctries)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	//fmt.Println(ctries)
	if err := json.NewEncoder(w).Encode(indicatordefs); err != nil {
		panic(err)
	}
}
