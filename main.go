package main

import (
	"html/template"
	"ip-web/infra"
	"ip-web/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var templates = template.Must(template.ParseFiles("./index.html", "./templates/dashboard.html", "./templates/mapa.html", "./templates/usuario.html", "./templates/head.html", "./templates/cadastrar-paciente.html", "./templates/pre-login.html", "./templates/teladelogin.html", "./templates/Sobre.html"))
var db = infra.CreateDatabaseConnection()

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/pacientes", pacientes)
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/mapa-interativo", mapa)
	http.HandleFunc("/usuario", usuario)
	http.HandleFunc("/paciente/cadastra", cadastrarPaciente)
	http.HandleFunc("/paciente/deleta", deletaPaciente)
	http.HandleFunc("/login", login)
	http.HandleFunc("/sobre", sobre)

	// Iniciar o servidor na porta 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
func sobre(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w,"Sobre.html", nil) 
		if err != nil{
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var dadosIncorretos bool = false

	senha := r.FormValue("senha")
	cpf := r.FormValue("Cpf")
	var Acs model.Acs

	db.Find(&Acs)

	if r.Method == "POST" {
		if Acs.Cpf == cpf && Acs.Senha == senha {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			dadosIncorretos = true
		}
	}
	
	err := templates.ExecuteTemplate(w, "teladelogin.html", dadosIncorretos)
	if err != nil {
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "pre-login.html", "a")
	if err != nil {
		return
	}
}

func pacientes(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	busca := strings.TrimSpace(r.Form.Get("busca"))

	var Pacientes []model.Paciente
	db.Find(&Pacientes)

	data := struct {
		Busca     string
		Pacientes []model.Paciente
	}{
		Busca:     busca,
		Pacientes: Pacientes,
	}

	db.Where("lower(nome) like lower(?)", "%"+data.Busca+"%").Order("nome").Find(&data.Pacientes)

	err = templates.Execute(w, data)
	if err != nil {
		return
	}
}

func dashboard(w http.ResponseWriter, _ *http.Request) {
	var acs model.Acs
	db.First(&acs)

	dataAtual := time.Now()
	hora, _, _ := dataAtual.Clock()
	ptDates := [][]string{{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sab"}, {"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}}
	date := ptDates[0][int(dataAtual.Weekday())] + ", " + strconv.Itoa(dataAtual.Day()) + " de " + ptDates[1][int(dataAtual.Month())-1] + " de " + strconv.Itoa(dataAtual.Year())

	greeting := ""
	if hora >= 4 && hora <= 12 {
		greeting = "Bom dia"
	} else if hora > 12 && hora <= 18 {
		greeting = "Boa tarde"
	} else {
		greeting = "Boa noite"
	}
	getFirstWord := func(name string) string {
		for i := range name {
			if string(name[i]) == " " {
				return name[0:i]
			}
		}
		return name
	}

	dashboardInfo := struct {
		AcsData      model.Acs
		Date         string
		Greeting     string
		AcsFirstName string
	}{
		AcsData:      acs,
		Date:         date,
		Greeting:     greeting,
		AcsFirstName: getFirstWord(acs.Nome),
	}

	err := templates.ExecuteTemplate(w, "dashboard.html", dashboardInfo)
	if err != nil {
		return
	}
}

func mapa(w http.ResponseWriter, _ *http.Request) {
	err := templates.ExecuteTemplate(w, "mapa.html", "a")
	if err != nil {
		return
	}
}

func usuario(w http.ResponseWriter, r *http.Request) {

	var acs model.Acs
	db.First(&acs)

	if r.Method == "POST" {
		acs.Nome = r.FormValue("nome")
		acs.Cpf = r.FormValue("CPF")
		acs.Cnes = r.FormValue("CNES")
		acs.Cns = r.FormValue("CNS")
		acs.Cbo = r.FormValue("CBO")
		acs.Ine = r.FormValue("INE")
		db.Save(&acs)
		http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
	}

	err := templates.ExecuteTemplate(w, "usuario.html", acs)
	if err != nil {
		return
	}
}

func cadastrarPaciente(w http.ResponseWriter, r *http.Request) {

	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")
	dataDeNascimento := r.FormValue("data_nasc")
	idade, _ := strconv.Atoi(r.FormValue("idade"))
	numero := r.FormValue("celular")
	sexo := r.FormValue("sexo")
	cep := r.FormValue("cep")
	estaFumante := r.Form.Has("esta_fumante")
	fazUsoDeAlcool := r.Form.Has("faz_uso_alcool")
	estaEmSituacaoDeRua := r.Form.Has("esta_situacao_de_rua")

	p := model.Paciente{
		Nome:              nome,
		Cpf:               cpf,
		DataNascimento:    dataDeNascimento,
		Idade:             idade,
		Telefone:          numero,
		Sexo:              sexo,
		Cep:               cep,
		EstaFumante:       estaFumante,
		FazUsoAlcool:      fazUsoDeAlcool,
		EstaSituacaoDeRua: estaEmSituacaoDeRua,
	}

	if sexo != "" {
		db.Create(&p)
	}

	err := templates.ExecuteTemplate(w, "cadastrar-paciente.html", nil)
	if err != nil {
		return
	}
}

func deletaPaciente(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	db.Delete(&model.Paciente{}, id)

	http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
}
