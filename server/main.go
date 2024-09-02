package main

import (
	"fmt"
	"net/http"
	"server/internal/api"

	"github.com/gorilla/mux"
)

func main() {

	var v = api.Server{}

	r := mux.NewRouter()

	r.HandleFunc("/User", v.GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/User", v.AnsverRequest).Methods(http.MethodPost)
	fmt.Println("сервер запущен")
	err := http.ListenAndServe("127.0.0.1:8080", r)
	//logger.Warn("сервер отключён")
	if err != nil {
		//logger.Error("сервер нe запустился")
		//log.Fatal(err)
		return
	}
}
