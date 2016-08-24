package controllers

import (
	//"fmt"
	"html/template"
	//"log"
	"net/http"

	//"encoding/json"
	"github.com/gorilla/context"
	//"github.com/thoas/stats"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	t, _ := template.ParseFiles("static/login.gtpl")
	t.Execute(w, nil)
	//fmt.Fprintf(w, "Hello (%v) %v\n", context.Get(r, "skey"),vars["id"])
}

func Index(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	t, err := template.ParseFiles("static/index.html") //.Delims("<<", ">>")
	if err != nil {
		//fmt.Fprintf(w, err)
		panic(err)
	}
	//fmt.Fprintf(w, "Hello \n")
	t.Execute(w, nil)

}

/*func Stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stats := stats_mw.Data()

	b, _ := json.Marshal(stats)

	w.Write(b)
}*/

func test(r *http.Request) {
	context.Set(r, "skey", "bar")
}
