// All the router handles should go here
package controller

import (
	"encoding/json"
	"fmt"
	"iban/internal/pkg/service"
	"log"
	"net/http"
)

//Error error struct for error handling...
type Error struct {
	Reason string `json:"reason,omitempty"`
}

//Request ...
type Request struct {
	IBAN string `json:"IBAN"`
}

//Response...
type Response struct {
	Valid bool   `json:"valid"`
	Error *Error `json:"error,omitempty"`
}

func handleResponse(w http.ResponseWriter, status int, response Response, e error) {
	w.Header().Set("Content-Type", "application/json")
	if e != nil {
		var error Error
		error.Reason = e.Error()
		response.Error = &error
	}
	resp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("error in parsing response, err: %v", err)
	}
	w.WriteHeader(status)
	w.Write(resp)

}

//IBANValidatorHandler...
func IBANValidatorHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	var res Response
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&req)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, res, err)
		return
	}
	if len(req.IBAN) < 5 {
		handleResponse(w, http.StatusBadRequest, res, fmt.Errorf("IBAN should be of length between 5 and 34"))
		return
	}
	isValid, err := service.CheckIBAN(req.IBAN)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, res, err)
		return
	}
	res.Valid = isValid
	handleResponse(w, http.StatusOK, res, nil)

}
