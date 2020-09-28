package main
import (
  
    "log"
    "net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"

)

type Company struct {
	CompId int `json:"id"`
	CompName string `json:"Cname"`
}
var company []Company
func main(){
	myRouter:=mux.NewRouter().StrictSlash(true)
	company=[]Company{
		Company{703078,"Infosys"},
		Company{703079,"CTS"},
	}
	
	
	myRouter.HandleFunc("/getdetails",GetAllDetails)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func GetAllDetails( w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode(company)

}