package infra

import (
	"ip-web/model"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=123 dbname=paciente port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("falha ao se conectar a DB ", err)
	}

	err = db.AutoMigrate(&model.Paciente{})
	if err != nil {
		log.Fatal("falha ao migrar database ", err)
	}

	return db
}