package goqrcode 

import (
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	// "strings"
)

func TestGenerateQRWithLogo(t *testing.T) {
	text := "https://pakarbi.github.io"
	logoPath := "logo_ulbi.png"
	outputPath := "codeqr.png"

	err := GenerateQRWithLogo(text, logoPath, outputPath)
	if err != nil {
		t.Errorf("error generating QR code with logo: %v", err)
	}

	// Optionally, add tests to check the output file or validate the generated QR code
	// For instance, check if the output file exists and has the expected properties.
	// Example: Validate existence, dimensions, and correctness of the generated QR code.
}

func TestGenerateQRFromEmail(t *testing.T) {
	Payload := Payload{Email: "test@example.com", Message: "Selamat Datang di PakArbi. Silakan lakukan verifikasi di email ULBI Anda."}
	reqBody, err := json.Marshal(Payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/generateQRFromEmail", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(generateQRFromEmail)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %d", rr.Code)
	}

	expected := "QR code generated successfully"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}