package model

type Paciente struct {
	ID                uint64 `gorm:"primary_key,autoIncrement"`
	Nome              string `json:"nome"`
	Cpf               string `json:"cpf"`
	DataNascimento    string `json:"data_nasc"`
	Idade             int    `json:"idade"`
	Telefone          string `json:"celular"`
	Sexo              string `json:"sexo"`
	Cep               string `json:"cep"`
	EstaFumante       bool   `json:"esta_fumante"`
	FazUsoAlcool      bool   `json:"faz_uso_alcool"`
	EstaSituacaoDeRua bool   `json:"esta_situacao_de_rua"`
}
