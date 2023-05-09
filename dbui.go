package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*
var embedded embed.FS

var templates *template.Template

func init() {
	var err error
	// Create a template with embedded HTML files
	// template names are defined in files, or usin filename
	templates, err = template.ParseFS(embedded, "templates/*")
	if err != nil {
		log.Fatal(err)
	}
}

func webui(port string) {

	// Define a handler functions
	http.HandleFunc("/", work)
	http.HandleFunc("/config", config)

	// Serve static files from the "static" directory
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	log.Printf("Listening on %s...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Println(err)
}

func work(w http.ResponseWriter, r *http.Request) {
	// Prepare the query
	rows, err := db.Query("SELECT lp, ticker, size FROM work")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Store the results in a slice
	var workitems []Work
	for rows.Next() {
		var work Work
		err := rows.Scan(&work.LP, &work.Ticker, &work.Size)
		if err != nil {
			log.Fatal(err)
		}
		workitems = append(workitems, work)
	}

	// Execute the template with the results
	err = templates.ExecuteTemplate(w, "work", workitems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func config(w http.ResponseWriter, r *http.Request) {
	// Prepare the query
	rows, err := db.Query("SELECT data_type, key, value FROM config")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Store the results in a slice
	var configs []Config
	for rows.Next() {
		var config Config
		err := rows.Scan(&config.Datatype, &config.Key, &config.Value)
		if err != nil {
			log.Fatal(err)
		}
		configs = append(configs, config)
	}

	// Execute the template with the results
	err = templates.ExecuteTemplate(w, "config", configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
