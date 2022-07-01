package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIBANValidatorHandlerError(t *testing.T) {
	handler := IBANValidatorHandler
	m, b := map[string]string{"IBAN": "567"}, new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	req, _ := http.NewRequest("POST", "/api/iban/validate", b)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)
	got := w.Code
	want := 400
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestIBANValidatorHandlerValid(t *testing.T) {
	handler := IBANValidatorHandler
	m, b := map[string]string{"IBAN": "DE89 3704 0044 0532 0130 00"}, new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	req, _ := http.NewRequest("POST", "/api/iban/validate", b)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)
	got := w.Code
	want := 200
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestIBANValidatorHandlerInValid(t *testing.T) {
	handler := IBANValidatorHandler
	m, b := map[string]string{"IBAN": "DE89 3794 9044 9532 0130 00"}, new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	req, _ := http.NewRequest("POST", "/api/iban/validate", b)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)
	got := w.Code
	want := 400
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
