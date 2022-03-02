package connection

import (
	"log"

	"github.com/custs-risk/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	Err error
)

func Connect() {
	DB, Err = gorm.Open("mysql", "root:@/custsrisk?charset=utf8&parseTime=True")

	if Err != nil {
		log.Println("Connection failed", Err)
	} else {
		log.Println("Server up and running")
	}

	DB.AutoMigrate(&structs.User{})
	DB.AutoMigrate(&structs.Risk_Profile{})
}
