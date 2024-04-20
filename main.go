package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"io/ioutil"
	"encoding/json"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/public/", logging(public()))
	mux.Handle("/", logging(index()))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("main: running simple server on port", port)
	log.Printf("localhost:%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: couldn't start simple server: %v\n", err)
	}
}

// logging is middleware for wrapping any handler we want to track response
// times for and to see what resources are requested.
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Since(start))
	})
}

var templates = template.Must(template.ParseFiles("./templates/index.html"))

// index is the handler responsible for rending the index page for the site.
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		type NasaApi struct {
			Copyright   	string
			Date 	    	string
			Explanation 	string
			Hdurl			string
			Media_type  	string
			Service_version string
			Title			string
			Url				string	   
		}

		var api NasaApi 

		err = json.Unmarshal(responseData, &api)
		if err != nil {
			log.Fatal(err)
		}
		b := struct {
			Title        template.HTML
			BusinessName string
			Slogan       string
			Nasa	     NasaApi
		}{
			Title:        template.HTML("Sismocabo"),
			BusinessName: "Sismocabo",
			Slogan:       "Sistema de Monitoramento do Câncer de Boca",
			Nasa: 		  api,
		}
		err = templates.ExecuteTemplate(w, "index", &b)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}

// func consumeApi() NasaApi {
	

// 	return api;
// } 
