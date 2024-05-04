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

var templates = template.Must(template.ParseFiles("./index.html", "./templates/dashboard.html", "./templates/mapa.html", "./templates/usuario.html", "./templates/head.html"))

func main() {
	infra.CreateConnection()

	http.HandleFunc("/", handler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/mapa-interativo", mapa)
	http.HandleFunc("/usuario", usuario)

	// Iniciar o servidor na porta 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	// busca := strings.TrimSpace(r.Form.Get("busca"))

	var Pacientes Pacientes

	jsonFile, _ := os.Open("data.json")
	byteJson, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteJson, &Pacientes)
	if err != nil {
		return
	}

	err = templates.Execute(w, Pacientes)
	if err != nil {
		return
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "dashboard.html", "a")
	if err != nil {
		return
	}
}

func mapa(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "mapa.html", "a")
	if err != nil {
		return
	}
}

func usuario(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "usuario.html", "a")
	if err != nil {
		return
	}
}

type Pacientes struct {
	Pacientes []model.Paciente `json:"pacientes"`
}
