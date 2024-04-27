package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseFiles("./templates/index.html", "./templates/dashboard.html", "./templates/mapa.html", "./templates/usuario.html"))

func main() {
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

	templates.Execute(w, Pacientes)
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

type Paciente struct {
	Nome              string `json:"nome"`
	Cpf               string `json:"cpf"`
	DataNascimento    string `json:"data_nasc"`
	Idade             int 	 `json:"idade"`
	Telefone          string `json:"celular"`
	Sexo              string `json:"sexo"`
	Cep               string `json:"cep"`
	EstaFumante       bool   `json:"esta_fumante"`
	FazUsoAlcool      bool   `json:"faz_uso_alcool"`
	EstaSituacaoDeRua bool   `json:"esta_situacao_de_rua"`
}

type Pacientes struct {
	Pacientes []Paciente `json:"pacientes"`
}