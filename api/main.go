package main

import (
	"api/database"
	"api/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	database.Create(os.Getenv("MYSQL_DATA"))
	defer database.Close()

	database.CreateInfo()

	http.Handle("/", http.FileServer(http.Dir("../web")))

	http.HandleFunc("/form", handler.FormHandler)
	http.HandleFunc("/query", handler.SearchHandler)
	http.HandleFunc("/data", handler.JsonHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
