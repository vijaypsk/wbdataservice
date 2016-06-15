package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/data", ShowData)
	r.HandleFunc("/data/{year}", GetIndicatorsForYear)
	r.HandleFunc("/data/{year}/{icode}", GetIndicatorDataForYear)
	r.HandleFunc("/definitions", GetIndicatorDefnitions)

	http.Handle("/", &MyServer{r})
	log.Fatal(http.ListenAndServe(":8000", nil))

	//log.Fatal(http.ListenAndServe(":8000", router))
}

//http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang
type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}
