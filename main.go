package main

import (
	"encoding/json"
	"html/template"
	"io"
	"ip-web/infra"
	"ip-web/model"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseGlob("./templates/*"))
var index = template.Must(template.ParseFiles("./index.html"));

func main() {
	infra.CreateConnection()

	http.HandleFunc("/", handler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/mapa-interativo", mapa)
	http.HandleFunc("/usuario", usuario)

	// Iniciar o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var Pacientes Pacientes

	jsonFile, _ := os.Open("data.json")
	byteJson, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteJson, &Pacientes)

	index.Execute(w, Pacientes)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dashboard.html", "a")
}

func mapa(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "mapa.html", "a")
}

func usuario(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "usuario.html", "a")
}

type Pacientes struct {
	Pacientes []model.Paciente `json:"pacientes"`
}