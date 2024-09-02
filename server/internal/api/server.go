package api

import (
	"encoding/json"
	"io"
	"net/http"
	"server/internal/db"
	"server/internal/domain"
)

type Server struct {
}

var v = Server{}

func (p Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(domain.Products)
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
	var date1 db.Client
	err = json.Unmarshal(jsong, &date1)
	if err != nil {
		//handleError(w, http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(date1.LoanCheck)
	if err != nil {
		//handleError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(data)

}
