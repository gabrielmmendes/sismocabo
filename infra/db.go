package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"ip-web/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", os.Getenv("host"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"), os.Getenv("port"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("falha ao se conectar a DB ", err)
	}

	err = db.AutoMigrate(&model.Paciente{})
	if err != nil {
		log.Fatal("falha ao migrar database ", err)
	}

	err = db.AutoMigrate(&model.Acs{})
	if err != nil {
		log.Fatal("falha ao migrar database ", err)
	}
	
	var Pacientes []model.Paciente
	db.Find(&Pacientes)
	if len(Pacientes) == 0 {
		Pacientes = jsonPacientesToList()
		db.Create(&Pacientes)
	}

	var Acss []model.Acs
	db.Find(&Acss)
	if len(Acss) == 0 {
		Acss = jsonAcsToList()
		db.Create(&Acss)
	}

	return db
}

type Pacientes struct {
	Pacientes []model.Paciente `json:"pacientes"`
}
func jsonPacientesToList() []model.Paciente {
	var Pacientes Pacientes

	jsonFile, _ := os.Open("data.json")
	byteJson, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteJson, &Pacientes)
	if err != nil {
		return nil
	}

	return Pacientes.Pacientes
}

type Acss struct {
	Acs []model.Acs `json:"acss"`
}
func jsonAcsToList() []model.Acs {
	var Acs Acss

	jsonFile, _ := os.Open("data.json")
	byteJson, _ := io.ReadAll(jsonFile)

	err := json.Unmarshal(byteJson, &Acs)
	if err != nil {
		return nil
	}
	
	return Acs.Acs
}
