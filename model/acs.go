package model

type Acs struct {
	ID    uint64 `gorm:"primary_key,autoIncrement"`
	Nome  string `json:"nome"`
	Cpf   string `json:"cpf"`
	Cns	  string `json:"cns"`
	Cbo   string `json:"cbo"`
	Ine   string `json:"ine"`
	Senha string `json:"senha"`
}