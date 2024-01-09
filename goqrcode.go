package goqrcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	// "image/draw"
	"net/http"
	"net/smtp"
	"os"

	"github.com/disintegration/imaging"
	"github.com/PakArbi/backparkir"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRCode(data Parkiran) (string, error) {
    qrData := prepareQRData(data) // Persiapkan data untuk QR code
    qrCode, err := qrcode.Encode(qrData, qrcode.Medium, 256)
    if err != nil {
        return "", err
    }
    return string(qrCode), nil
}

func prepareQRData(data Parkiran) string {
    qrData := fmt.Sprintf("ID: %s\nParkiranID: %d\nNama: %s\nNPM: %s\nProdi: %s\nNama Kendaraan: %s\nNomor Kendaraan: %s\nJenis Kendaraan: %s",
        data.ID.Hex(), data.ParkiranId, data.Nama, data.NPM, data.Prodi, data.NamaKendaraan, data.NomorKendaraan, data.JenisKendaraan)
    return qrData
}

func SendEmail(to, subject, body string) error {
    from := "your_email@gmail.com"
    password := "your_password"
    
    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
    err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
    if err != nil {
        return err
    }
    return nil
}

func GenerateHandler(MONGOCONNSTRINGENV, dbname, collectionname string, w http.ResponseWriter, r *http.Request) string {
    if r.Method == "POST" {
        var parkiranData Parkiran
        err := json.NewDecoder(r.Body).Decode(&parkiranData)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Generate QR Code
        qrData := prepareQRData(parkiranData)
        qrCode, err := GenerateQRCode(qrData)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Simpan hasil scan QR ke MongoDB
        qrScanData := QRScan{
            QR:      qrCode,
            Status:  "scanned",
            Message: "QR code scanned and stored",
        }

        err = SaveQRScanResult(qrScanData)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

         // Menerima informasi dari alert
    	alertInfo := getAlertInfo(r) // Misalnya, mendapatkan informasi dari request

    	// Lakukan sesuatu dengan informasi yang diterima dari alert
    	err = HandleAlertInfo(alertInfo)
    	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    	}

        // Kirim respons ke pengguna
        response := Notifikasi{
            Status:  http.StatusOK,
            Message: "Data berhasil diproses",
            Data:    parkiranData,
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
}

func GenerateQRCodeWithLogo(DataParkir backparkir.Parkiran) error {
	// Convert struct to JSON
	dataJSON, err := json.Marshal(DataParkir)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Generate QR code
	qrCode, err := qrcode.Encode(string(dataJSON), qrcode.Medium, 256)
	if err != nil {
		return fmt.Errorf("failed to generate QR code: %v", err)
	}

	// Create an image from the QR code
	qrImage, err := imaging.Decode(bytes.NewReader(qrCode))
	if err != nil {
		return fmt.Errorf("failed to decode QR code image: %v", err)
	}

	// Open the ULBI logo file
	logoFile, err := os.Open("logo_ulbi.png") // Replace with your ULBI logo file path
	if err != nil {
		return fmt.Errorf("failed to open logo file: %v", err)
	}
	defer logoFile.Close()

	// Decode the ULBI logo
	logo, _, err := image.Decode(logoFile)
	if err != nil {
		return fmt.Errorf("failed to decode logo image: %v", err)
	}

	// Resize the logo to fit within the QR code
	resizedLogo := imaging.Resize(logo, 80, 0, imaging.Lanczos)

	// Calculate position to overlay the logo on the QR code
	x := (qrImage.Bounds().Dx() - resizedLogo.Bounds().Dx()) / 2
	y := (qrImage.Bounds().Dy() - resizedLogo.Bounds().Dy()) / 2

	// Draw the logo onto the QR code
	result := imaging.Overlay(qrImage, resizedLogo, image.Pt(x, y), 1.0)

	// Save the final QR code with logo
	outFile, err := os.Create("qrcode.png") // Replace with desired output file name
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	// Encode the final image into the output file
	err = imaging.Encode(outFile, result, imaging.PNG)
	if err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}

	return nil
}

func GCFGenerateQRCode(w http.ResponseWriter, r *http.Request) {
	var DataParkir backparkir.Parkiran
	err := json.NewDecoder(r.Body).Decode(&DataParkir)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}

	// Generate QR code with logo
	err = GenerateQRCode(DataParkir)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate QR code with logo: %v", err), http.StatusInternalServerError)
		return
	}

	// Set response content type
	w.Header().Set("Content-Type", "image/png")

	// Open and serve the QR code image with logo
	http.ServeFile(w, r, "qrcode.png")
}





