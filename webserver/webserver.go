package webserver

import (
	"fmt"
	//"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"bookmarks_service/controllers"

	//"encoding/json"
	//"github.com/gorilla/context"
	"github.com/gorilla/mux"
	//"github.com/thoas/stats"
	"github.com/urfave/negroni"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

func DB(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
	dbs, err := session.DatabaseNames()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, strings.Join(dbs, ","))

	//fmt.Fprintf(w, "Hello ")
}

func Start() {
	//stats_mw := stats.New()

	var err error
	session, err = mgo.Dial("185.45.193.231/my_data")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	mux := mux.NewRouter()

	mux.HandleFunc("/login", controllers.Login).Methods("GET")
	mux.HandleFunc("/db", DB).Methods("GET")
	//mux.HandleFunc("/stats", controllers.Stats).Methods("GET")
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	mux.HandleFunc("/", controllers.Index).Methods("GET")
	//n := negroni.Classic()
	n := negroni.New()

	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	//n.Use(stats_mw)
	n.UseHandler(mux)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
	//http.ListenAndServe(":3000", n)
}
