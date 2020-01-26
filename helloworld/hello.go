package helloworld

import (
	"fmt"
	"net/http"
)

//Hello write a simple - Docker - Hello World !
func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Docker - Hello World !")

}