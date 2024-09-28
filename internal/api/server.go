package api

import (
	"encoding/json"
	"io"
	"net/http"

	"server/internal/domain"
)

type Server struct {
	DateBase domain.Client
}

func (p Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(domain.Products) //(domain.Products)????
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)
}
func (p Server) AnsverRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsong, err := io.ReadAll(r.Body)
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	var DateBase2 domain.Client
	//var n = domain.Client(DateBase2)
	err = json.Unmarshal(jsong, &DateBase2) //data1) ???
	//fmt.Println(DateBase2)
	if err != nil {
		//handleError(w, http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(domain.LoanCheck(DateBase2))
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)

}
