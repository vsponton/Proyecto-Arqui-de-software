package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"

	"cursos-ucc/clients/user"
	"cursos-ucc/clients/course"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "test"  //Nombre de la base de datos local de ustedes
	DBUser := "root"      //usuario de la base de datos, habitualmente root
	DBPass := "root"            //password del root en la instalacion
	DBHost := "localhost" //host de la base de datos. hbitualmente 127.0.0.1
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	courseClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	// db.AutoMigrate(&model.Barrio{})
	//db.AutoMigrate(&model.Sensor{})
	//db.AutoMigrate(&model.Medicion{})

	log.Info("Finishing Migration Database Tables")
}
