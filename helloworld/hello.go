package helloworld

import (
	"fmt"
	"github.com/idasilva/go-dockerfile/db"
	"log"
	"net/http"
)

//Hello write a simple - Docker - Hello World !
func Hello(w http.ResponseWriter, r *http.Request){
	err := db.Conn()
	if err != nil {
		fmt.Fprint(w,"error connect to database:",err)

	}

	db.Db.AutoMigrate(&db.User{})
	db.Db.Create(&db.User{
		ID:       1,
		Username: "idasilva1",
	})
	ida := db.User{}
	db.Db.Find(&ida)
	fmt.Fprint(w, ida)


	log.Printf("Connection to database sucessfull !!")

	fmt.Fprint(w,"Docker - Hello World !---")

}