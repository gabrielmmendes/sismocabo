package main

import (
    "net/http"
)

func main() {
    // Configuração do servidor para servir arquivos estáticos (HTML, CSS, JS, imagens, etc.)
    fs := http.FileServer(http.Dir("./"))
    http.Handle("/", fs)

    // Iniciar o servidor na porta 8080
    http.ListenAndServe(":8080", nil)
}