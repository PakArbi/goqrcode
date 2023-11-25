package goqrcode

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	// "net/http"
	// "io/ioutil"
	// "encoding/base64"
	"encoding/json"

	qrcode "github.com/skip2/go-qrcode"
	// qrpa "github.com/PakArbi/model/goqrcode"
	"github.com/nfnt/resize"
)

// func GenerateQRWithLogo(text, logoFilename, outputPath string) error {
// 	// Generate QR code
// 	qrCode, err := qrcode.Encode(text, qrcode.Medium, 256)
// 	if err != nil {
// 		return fmt.Errorf("error generating QR code: %v", err)
// 	}

// 	// Open logo file from 'img' folder
// 	imgFolder := "img/logo_ulbi.png"
// 	logoPath := filepath.Join(imgFolder, logoFilename)
// 	logoFile, err := os.Open(logoPath)
// 	if err != nil {
// 		return fmt.Errorf("error opening logo file: %v", err)
// 	}
// 	defer logoFile.Close()

// 	// Decode logo image
// 	logo, _, err := image.Decode(logoFile)
// 	if err != nil {
// 		return fmt.Errorf("error decoding logo image: %v", err)
// 	}

// 	// Decode QR code image
// 	qrImage, _, err := image.Decode(bytes.NewReader(qrCode))
// 	if err != nil {
// 		return fmt.Errorf("error decoding QR code image: %v", err)
// 	}

// 	// Create an RGBA image to draw QR code and logo
// 	rgba := image.NewRGBA(qrImage.Bounds())
// 	draw.Draw(rgba, qrImage.Bounds(), qrImage, image.Point{}, draw.Over)

// 	// Resize the logo to fit within the QR code
// 	resizedLogo := resize.Resize(80, 0, logo, resize.Lanczos3)

// 	// Calculate position to overlay the logo on the QR code
// 	x := (qrImage.Bounds().Dx() - resizedLogo.Bounds().Dx()) / 2
// 	y := (qrImage.Bounds().Dy() - resizedLogo.Bounds().Dy()) / 2

// 	// Draw the logo onto the QR code
// 	draw.Draw(rgba, resizedLogo.Bounds().Add(image.Point{x, y}), resizedLogo, image.Point{}, draw.Over)

// 	// Save the final QR code with logo
// 	outFile, err := os.Create(outputPath)
// 	if err != nil {
// 		return fmt.Errorf("error creating output file: %v", err)
// 	}
// 	defer outFile.Close()

// 	// Encode the final image into the output file
// 	err = png.Encode(outFile, rgba)
// 	if err != nil {
// 		return fmt.Errorf("error encoding image: %v", err)
// 	}

// 	return nil
// }

func GenerateQRCode(formData FormData) error {
    // Convert struct to JSON
    dataJSON, err := json.Marshal(formData)
    if err != nil {
        return fmt.Errorf("failed to marshal JSON: %v", err)
    }

    // Generate QR code
    qrCode, err := qrcode.Encode(string(dataJSON), qrcode.Medium, 256)
    if err != nil {
        return fmt.Errorf("failed to generate QR code: %v", err)
    }

    // Open ULBI logo file
    logoFile, err := os.Open("./img/logo_ulbi.png") // Replace with your ULBI logo file path
    if err != nil {
        return fmt.Errorf("failed to open logo file: %v", err)
    }
    defer logoFile.Close()

    // Decode ULBI logo
    logo, _, err := image.Decode(logoFile)
    if err != nil {
        return fmt.Errorf("failed to decode logo image: %v", err)
    }

    // Decode QR code image
    qrImage, _, err := image.Decode(bytes.NewReader(qrCode))
    if err != nil {
        return fmt.Errorf("failed to decode QR code image: %v", err)
    }

    // Create an RGBA image to draw QR code and logo
    rgba := image.NewRGBA(qrImage.Bounds())
    draw.Draw(rgba, qrImage.Bounds(), qrImage, image.Point{}, draw.Over)

    // Resize the logo to fit within the QR code
    resizedLogo := resize.Resize(80, 0, logo, resize.Lanczos3)

    // Calculate position to overlay the logo on the QR code
    x := (qrImage.Bounds().Dx() - resizedLogo.Bounds().Dx()) / 2
    y := (qrImage.Bounds().Dy() - resizedLogo.Bounds().Dy()) / 2

    // Draw the logo onto the QR code
    draw.Draw(rgba, resizedLogo.Bounds().Add(image.Point{x, y}), resizedLogo, image.Point{}, draw.Over)

    // Save the final QR code with logo
    outFile, err := os.Create("./img/qrcode.png") // Replace with desired output file name
    if err != nil {
        return fmt.Errorf("failed to create output file: %v", err)
    }
    defer outFile.Close()

    // Encode the final image into the output file
    err = png.Encode(outFile, rgba)
    if err != nil {
        return fmt.Errorf("failed to encode image: %v", err)
    }

    return nil
}

// func generateQRFromEmail(w http.ResponseWriter, r *http.Request, db Database) {
// 	// Read request body data
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Failed to read request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Parse JSON from the request body
// 	var emailData EmailData
// 	err = json.Unmarshal(reqBody, &emailData)
// 	if err != nil {
// 		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Generate QR code from the email address
// 	err = GenerateQRWithLogo(emailData.Email, "logo_ulbi.png", "codeqr.png")
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to generate QR code: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	payload := Payload{
// 		Email:   emailData.Email,
// 		Message: "Selamat Anda berhasil verifikasi",
// 	}

// 	// Insert payload data into the database
// 	err = db.InsertPayload(payload)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to save user: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("QR code generated successfully"))
// }

//handler