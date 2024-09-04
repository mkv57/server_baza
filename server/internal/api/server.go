package api

import (
	"encoding/json"
	"io"
	"net/http"

	"server/internal/domain"
)

type Server struct {
	DateBase domain.Client
	DateProd domain.Loan
}

func (p Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(p.DateProd) //(domain.Products)????
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
	//var date1 db.Client ???
	err = json.Unmarshal(jsong, &p.DateBase) //data1) ???
	if err != nil {
		//handleError(w, http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(p.DateBase.LoanCheck)
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)

}
