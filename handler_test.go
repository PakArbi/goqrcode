package goqrcode 

import (
	"testing"
	"strings"
	// "bytes"
	"net/http"
	"net/http/httptest"
	// "encoding/json"
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
	// Create a sample payload for testing
	payload := `{"email": "test@example.com", "message": "Hello, World!"}`

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/generateQRFromEmail", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to capture the handler response
	rr := httptest.NewRecorder()

		// Mock database instance
		mockDB := &MockDatabase{} // Use pointer to MockDatabase

		// Call the handler function
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			generateQRFromEmail(w, r, mockDB)
		})
		
	// Serve the HTTP request and record the response
	handler.ServeHTTP(rr, req)

	// Check if the status code is as expected
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %d", rr.Code)
	}

	// Check the response body if needed
	expected := "QR code generated successfully"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}