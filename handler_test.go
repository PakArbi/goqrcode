package goqrcode 

import (
	"testing"
	// "bytes"
	"os"
	// "fmt"

	// "net/http"
	// "net/http/httptest"
	"encoding/json"
	"strings"
)

// func TestGenerateQRWithLogo_EmailNPM(t *testing.T) {
// 	email := "npm@std.ulbi.ac.id"
// 	outputPath := "output_qr.png" // Set your desired output path

// 	err := GenerateQRWithLogo(email, "logo_ulbi.png", outputPath)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	// Check if the output file exists
// 	_, err = os.Stat(outputPath)
// 	if os.IsNotExist(err) {
// 		t.Errorf("Output file not created: %v", err)
// 	}

// 	// Clean up: Remove the created output file
// 	err = os.Remove(outputPath)
// 	if err != nil {
// 		fmt.Println("Error removing output file:", err)
// 	}
// }

func TestGenerateQRCode(t *testing.T) {
	formData := FormData{
		NamaLengkap:    "Farhan Rizki Maulana",
		NPM:            "1214020",
		Jurusan:        "D4 Teknik Informatika",
		NamaKendaraan:  "Supra X 125",
		NomorKendaraan: "F 1234 NR",
		JenisKendaraan: "Motor",
	}

	err := GenerateQRCode(formData)
	if err != nil {
		t.Errorf("Failed to generate QR code: %v", err)
	}

	// Check if QR code file exists
	if _, err := os.Stat("qrcode.png"); os.IsNotExist(err) {
		t.Errorf("QR code file does not exist: %v", err)
	}

	// Check if JSON data is generated correctly
	expectedJSON := `{"namalengkap":"Farhan Rizki Maulana",
	"npm":"1214020","jurusan":"D4 Teknik Informatika",
	"namakendaraan":"Supra X 125",
	"nomorkendaraan":"F 1234 NR",
	"jeniskendaraan":"Motor"}`
	
	// Convert struct to JSON string
	dataJSON, err := json.Marshal(formData)
	if err != nil {
		t.Errorf("Failed to marshal JSON: %v", err)
	}

	generatedJSON := string(dataJSON)

	// Remove extra spaces and new lines from both JSON strings
	expectedJSON = strings.ReplaceAll(expectedJSON, " ", "")
	expectedJSON = strings.ReplaceAll(expectedJSON, "\n", "")
	generatedJSON = strings.ReplaceAll(generatedJSON, " ", "")
	generatedJSON = strings.ReplaceAll(generatedJSON, "\n", "")

	if expectedJSON != generatedJSON {
		t.Errorf("Incorrect JSON data generated")
	}
}
