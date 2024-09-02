package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/internal/db"
	"server/internal/domain"

	"github.com/gorilla/mux"
	//"server/structs"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(domain.Products)
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)
}
func AnsverRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsong, err := io.ReadAll(r.Body)
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	var date1 db.Client
	err = json.Unmarshal(jsong, &date1)
	if err != nil {
		//handleError(w, http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(db.LoanCheck(date1))
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/User", GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/User", AnsverRequest).Methods(http.MethodPost)
	fmt.Println("сервер запущен")
	err := http.ListenAndServe("127.0.0.1:8080", r)
	//logger.Warn("сервер отключён")
	if err != nil {
		//logger.Error("сервер нe запустился")
		//log.Fatal(err)
		return
	}
}

//var dataClient = Client{}

/*fmt.Println("введите ваше имя")
fmt.Scanln(&dataClient.Name)
fmt.Println("введите ваш возраст")
fmt.Scanln(&dataClient.Age)
fmt.Println("введите вашу з/п в руб.")
fmt.Scanln(&dataClient.Wage)

fmt.Println("Уважаемый ", dataClient.Name)
LoanCheck(dataClient)
*/
