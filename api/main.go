package main

import (
	"api/database"
	"fmt"
	"log"
	"net/http"
	"os"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "Invalid Method")
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	database.InsertInfo(r.FormValue("fname"), r.FormValue(("lname")))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Invalid Method")
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	database.GetName(r.FormValue("fullname"), w)
}

func main() {
	database.Create(os.Getenv("MYSQL_DATA"))
	defer database.Close()

	database.CreateInfo()

	http.Handle("/", http.FileServer(http.Dir("../web")))
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/query", searchHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
