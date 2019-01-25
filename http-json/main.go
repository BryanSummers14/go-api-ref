package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type profile struct {
	Name    string   `json:"name,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}

type threeformbuzz struct {
	ID            int64  `json:"id,omitempty"`
	Date          string `json:"date,omitempty"`
	Contact1      int64  `json:"contact_1,omitempty"`
	Contact2      int64  `json:"contact_2,omitempty"`
	SocialHeaders string `json:"social_headers,omitempty"`
	SocialTip     string `json:"social_tip,omitempty"`
	Generated     int64  `json:"generated,omitempty"`
}

func main() {
	r := mux.NewRouter()
	fmt.Println("Server is listening on port: 3000")
	go r.HandleFunc("/", foo).Methods("GET")
	go r.HandleFunc("/test-db", dbTest).Methods("GET")
	//http.HandleFunc("/", foo)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := profile{Name: "Alex", Hobbies: []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func dbTest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "3form:3formusa@/3formusa")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT * FROM 3formusa.3formbuzz")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var buzz []threeformbuzz

	results, err := stmtOut.Query()

	for results.Next() {
		var (
			ID            int64
			Date          sql.NullString
			Contact1      sql.NullInt64
			Contact2      sql.NullInt64
			SocialHeaders sql.NullString
			SocialTip     sql.NullString
			Generated     sql.NullInt64
		)
		err = results.Scan(&ID, &Date, &Contact1, &Contact2, &SocialHeaders, &SocialTip, &Generated)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		buzzRow := threeformbuzz{ID: ID, Date: Date.String, Contact1: Contact1.Int64, Contact2: Contact2.Int64, SocialHeaders: SocialHeaders.String, SocialTip: SocialTip.String, Generated: Generated.Int64}
		buzz = append(buzz, buzzRow)
	}

	js, err := json.Marshal(buzz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
