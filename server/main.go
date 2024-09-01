package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	//"server/structs"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal("hello world")
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/User", GetProduct).Methods(http.MethodGet)
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
