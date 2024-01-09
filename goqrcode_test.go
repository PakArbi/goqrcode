package goqrcode 

import (
	// "bytes"
	// "net/http"
	// "net/http/httptest"
	"testing"
	// "strings"
	"fmt"
	"os"
	"io/ioutil"

	// "encoding/json"

	"github.com/PakArbi/backparkir"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGenerateQRCode(t *testing.T) {
    dataParkir := backparkir.Parkiran{
        ParkiranId:      3,
        Nama:            "M Faisal A",
        NPM:             "1214000",
        Prodi:           "Teknik Informatika",
        NamaKendaraan:   "Honda",
        NomorKendaraan:  "D 1234 CD",
        JenisKendaraan:  "Motor",
    }

    err := GenerateQRCodeWithLogo(dataParkir)
    if err != nil {
        t.Errorf("Failed to generate QR code: %v", err)
    }

    // Check if QR code file exists
    if _, err = os.Stat("qrcode.png"); os.IsNotExist(err) {
        t.Errorf("QR code file does not exist: %v", err)
    } else {
        t.Logf("QR code generated successfully")
    }

    // Generate expected JSON with an auto-incremented ID based on NPM
    expectedJSON := fmt.Sprintf(`{"ID":"D3%s","parkiranid":3,"nama":"M Faisal A","npm":"1214000","prodi":"Teknik Informatika","namakendaraan":"Honda","nomorkendaraan":"D 1234 CD","jeniskendaraan":"Motor"}`, dataParkir.NPM[len(dataParkir.NPM)-4:])

    // Read the generated JSON file
    jsonData, err := ioutil.ReadFile("qrcode.png") // Change to the generated JSON file path
    if err != nil {
        t.Errorf("Failed to read generated JSON file: %v", err)
        return
    }

    // Convert JSON bytes to string
    jsonString := string(jsonData)

    // Validate JSON data
    if jsonString != expectedJSON {
        t.Errorf("Incorrect JSON data generated")
    } else {
        t.Logf("Generated JSON data matches the expected format")
    }

	 // If all checks passed, log a PASS message
	 t.Logf("OK PASS")
}




// func TestGCFGenerate(t *testing.T) {
//     // Prepare a sample request body
//     requestBody := `{
//         "_id": "some_id",
//         "parkiranid": 123,
//         "nama": "John Doe",
//         "npm": "12345",
//         "prodi": "Computer Science",
//         "namakendaraan": "Car",
//         "nomorkendaraan": "ABC123",
//         "jeniskendaraan": "Sedan"
//     }`

//     // Create a sample request
//     req := httptest.NewRequest("POST", "/codeqr", strings.NewReader(requestBody))

//     // Create a ResponseRecorder to capture the response
//     rr := httptest.NewRecorder()

//     // Call the handler directly (GCFGenerate)
//     GCFGenerateCodeQR(rr, req)

//    // Check the response status code
//    if status := rr.Code; status != http.StatusOK {
// 	t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// }

// // Check the Content-Type header
// contentType := rr.Header().Get("Content-Type")
// if contentType != "application/json" {
// 	t.Errorf("handler returned wrong content type: got %v want application/json", contentType)
// }

// // Check the response body
// var response struct {
// 	Message string `json:"message"`
// }
// err := json.NewDecoder(rr.Body).Decode(&response)
// if err != nil {
// 	t.Errorf("failed to decode response body: %v", err)
// }
// expectedResponseBody := "QR code generated successfully"
// if response.Message != expectedResponseBody {
// 	t.Errorf("handler returned unexpected body: got %v want %v", response.Message, expectedResponseBody)
// }
// }




//test GCF

// func TestGCFGenerateQRCode(t *testing.T) {
// 	// Buat data dummy untuk dijadikan input
// 	DataParkir := backparkir.Parkiran{
// 		ParkiranId:      3,
// 		Nama:            "M Faisal A",
// 		NPM:             "1214000",
// 		Prodi:           "Teknik Informatika",
// 		NamaKendaraan:   "Honda",
// 		NomorKendaraan:  "D 1234 CD",
// 		JenisKendaraan:  "Motor",
// 	}

// 	// Marshal data dummy ke JSON
// 	dataJSON, err := json.Marshal(DataParkir)
// 	if err != nil {
// 		t.Fatalf("Failed to marshal JSON: %v", err)
// 	}

// 	// Buat HTTP request dengan JSON data sebagai body
// 	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(dataJSON))
// 	if err != nil {
// 		t.Fatalf("Failed to create request: %v", err)
// 	}

// 	// Inisialisasi ResponseRecorder untuk merekam response
// 	rr := httptest.NewRecorder()

// 	// Panggil GCFGenerateQRCode handler function
// 	GCFGenerateQRCode(rr, req)

// 	// Cek status code
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	// Cek tipe konten response
// 	contentType := rr.Header().Get("Content-Type")
// 	if contentType != "image/png" {
// 		t.Errorf("Handler returned wrong content type: got %v want image/png", contentType)
// 	}
// 	// Periksa konten response lebih spesifik jika diperlukan
// }

