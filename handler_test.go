package goqrcode 

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestGenerateQRWithLogo(t *testing.T) {
	text := "https://ulbi.ac.id"
	logoPath := "logo_ulbi.png"
	outputPath := "qrcode_with_logo.png"

	err := GenerateQRWithLogo(text, logoPath, outputPath)
	if err != nil {
		t.Errorf("error generating QR code with logo: %v", err)
	}

	// Optionally, add tests to check the output file or validate the generated QR code
	// For instance, check if the output file exists and has the expected properties.
	// Example: Validate existence, dimensions, and correctness of the generated QR code.
}

func TestGenerateQRFromEmail(t *testing.T) {
	payload := `{"email": "test@example.com"}`
	req, err := http.NewRequest("POST", "/generateQRFromEmail", strings.NewReader(payload))
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