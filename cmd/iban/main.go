// This is a golang rest application to cheak if a IBAN number is valid
package main

import (
	"iban/internal/pkg/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/iban/validate", controller.IBANValidatorHandler)
	log.Fatal(http.ListenAndServe(":8080", nil).Error())
}
