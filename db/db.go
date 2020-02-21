package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Db stores the database connection
var Db *gorm.DB

type User struct {
	ID int
	Username string
}


//Conn connection to database
func Conn()(err error){
	Db, err = gorm.Open("postgres","host=db port=5432 user=idasilva dbname=gorm password=idasilva sslmode=disable")


	if err != nil {
		return err
	}
	return

}