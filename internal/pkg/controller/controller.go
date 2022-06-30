// All the router handles should go here
package controller

import (
	"encoding/json"
	"iban/internal/pkg/service"
	"log"
	"net/http"
)

//Error error struct for error handling...
type Error struct {
	Error string
}

//Request ...
type Request struct {
	IBAN string `json:"IBAN"`
}

//Response...
type Response struct {
	Valid bool `json:"valid"`
}

//handleErrorResponse used to generate a error response
func handleErrorResponse(w http.ResponseWriter, status int, message string) {
	response := Error{}
	response.Error = message
	resp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("error in parsing response, err: %v", err)
	}
	w.WriteHeader(status)
	w.Write(resp)
}
func handleSuccessResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("error in parsing response, err: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

//IBANValidatorHandler...
func IBANValidatorHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req Request
	err := decoder.Decode(&req)
	if err != nil {
		handleErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	var res Response
	if len(req.IBAN) < 5 {
		handleErrorResponse(w, http.StatusBadRequest, "IBAN should be of length between 5 and 34")
		return
	}
	isValid, err := service.CheckIBAN(req.IBAN)
	if err != nil {
		handleErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	res.Valid = isValid
	handleSuccessResponse(w, res)

}
