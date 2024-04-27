package main

import (
    "net/http"
    "html/template"
)

func main() {
    http.HandleFunc("/", handler)

    // Iniciar o servidor na porta 8080
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("index.html")

    p := Pacientes{
        Itens: []Paciente{
            {
                Nome:           "Diogo SIlva Lima",
                Cpf:            "867.420.587-90",
            },
            {
                Nome:           "Diogo SIlva Lima",
                Cpf:            "867.420.587-90",
            },
        },
    }

    tmpl.Execute(w, p)
}

type Pacientes struct {
    Itens []Paciente
}

type Paciente struct {
    Nome           string
    Cpf            string
    DataNascimento string
    Telefone       string
    Sexo           string
    EstaFumante    bool
    FazUsoAlcool   bool 
    SituacaoDeRua  bool
}