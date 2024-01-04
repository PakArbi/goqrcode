package goqrcode 

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"os"

	"github.com/stretchr/testify/assert"
	"github.com/PakArbi/backparkir"
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
	formData := backparkir.Parkiran{
		ParkiranId:      1,
		Nama:            "M Faisal A",
		NPM:             "1214000",
		Prodi:           "Teknik Informatika",
		NamaKendaraan:   "Honda",
		NomorKendaraan:  "D 1234 CD",
		JenisKendaraan:  "Motor",
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
	dataJSON, _ := json.Marshal(formData)

	expectedJSON := `{"ParkiranId":1,"Nama":"M Faisal A","NPM":"1214000","Prodi":"Teknik Informatika","NamaKendaraan":"Honda","NomorKendaraan":"D 1234 CD","JenisKendaraan":"Motor"}`

	// Validate JSON data
	if string(dataJSON) != expectedJSON {
		t.Errorf("Incorrect JSON data generated")
	}
}

func TestGCFGenerateQRCode(t *testing.T) {
	// Buat data dummy untuk dijadikan input
	formData := backparkir.Parkiran{
		ParkiranId:      1,
		Nama:            "M Faisal A",
		NPM:             "1214000",
		Prodi:           "Teknik Informatika",
		NamaKendaraan:   "Honda",
		NomorKendaraan:  "D 1234 CD",
		JenisKendaraan:  "Motor",
	}

	// Marshal data dummy ke JSON
	dataJSON, err := json.Marshal(formData)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Buat HTTP request dengan JSON data sebagai body
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Inisialisasi ResponseRecorder untuk merekam response
	rr := httptest.NewRecorder()

	// Panggil GCFGenerateQRCode handler function
	GCFGenerateQRCode(rr, req)

	// Cek status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Cek tipe konten response
	contentType := rr.Header().Get("Content-Type")
	if contentType != "image/png" {
		t.Errorf("Handler returned wrong content type: got %v want image/png", contentType)
	}
	// Periksa konten response lebih spesifik jika diperlukan
}

