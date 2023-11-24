package goqrcode 

import (
	"testing"
	// "bytes"
	"os"
	"fmt"

	// "net/http"
	// "net/http/httptest"
	// "encoding/json"
	// "strings"
)

func TestGenerateQRWithLogo_EmailNPM(t *testing.T) {
	email := "npm@std.ulbi.ac.id"
	outputPath := "output_qr.png" // Set your desired output path

	err := GenerateQRWithLogo(email, "logo_ulbi.png", outputPath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output file exists
	_, err = os.Stat(outputPath)
	if os.IsNotExist(err) {
		t.Errorf("Output file not created: %v", err)
	}

	// Clean up: Remove the created output file
	err = os.Remove(outputPath)
	if err != nil {
		fmt.Println("Error removing output file:", err)
	}
}

// TestInsertPayloadIntoMongo tests the insertion of payload into MongoDB
func TestInsertPayloadIntoMongo(t *testing.T) {
	// Set up your test payload
	payload := Payload{
		Email:   "test@example.com",
		Message: "Test message",
	}

	// Mock MongoDB instance for testing
	mockDB := &MockDatabase{}

	// Call the InsertPayload function with the mock database
	err := InsertPayload(payload)
	if err != nil {
		t.Errorf("Failed to insert payload into MongoDB: %v", err)
	}

	// Check if the InsertPayloadFunc in the mock database was called
	if !mockDB.InsertPayloadFuncCalled {
		t.Error("InsertPayloadFunc not called")
	}
}

// TestSendVerificationEmail tests the email verification functionality
func TestSendVerificationEmail(t *testing.T) {
	// Set up test email address
	email := "test@example.com"

	// Mock email sender for testing
	mockEmailSender := &MockEmailSender{}

	// Call the SendVerificationEmail function with the mock email sender
	err := SendVerificationEmail(email)
	if err != nil {
		t.Errorf("Failed to send verification email: %v", err)
	}

	// Check if the SendVerificationEmail function in the mock email sender was called
	if !mockEmailSender.SendVerificationEmailFuncCalled {
		t.Error("SendVerificationEmailFunc not called")
	}
}

// func TestVerifyEmail(t *testing.T) {
// 	// Mock the email verification process

// 	// Mock request body
// 	email := "test@example.com"
// 	requestBody, _ := json.Marshal(EmailData{Email: email})

// 	// Create a mock HTTP request with the request body
// 	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a mock response recorder
// 	rr := httptest.NewRecorder()

// 	// Mock Database instance
// 	db := &MockDB{}

// 	// Call the handler function with the mock request, response recorder, and mock DB
// 	generateQRFromEmail(rr, req, db)

// 	// Check the HTTP status code
// 	if rr.Code != http.StatusOK {
// 		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
// 	}

// 	// Add assertions to check if the email verification was triggered or mocked appropriately
// 	// For example, check if the email was sent by verifying specific behaviors or functions were called
// }

// func TestGenerateQRFromEmail(t *testing.T) {
// 	// Mock request body
// 	email := "test@example.com"
// 	requestBody, _ := json.Marshal(EmailData{Email: email})

// 	// Create a mock HTTP request with the request body
// 	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a mock response recorder
// 	rr := httptest.NewRecorder()

// 	// Mock Database instance
// 	db := &MockDB{}

// 	// Call the handler function with the mock request, response recorder, and mock DB
// 	generateQRFromEmail(rr, req, db)

// 	// Check the HTTP status code
// 	if rr.Code != http.StatusOK {
// 		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
// 	}

// 	// Check the response body
// 	expectedBody := "QR code generated successfully"
// 	if rr.Body.String() != expectedBody {
// 		t.Errorf("Handler returned unexpected body: got %s, want %s", rr.Body.String(), expectedBody)
// 	}

// 	// Add more assertions to validate further functionality (e.g., check if the QR code is generated, email sent, etc.)
// }
// ghfg
