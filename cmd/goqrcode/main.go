package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	qrcode "github.com/skip2/go-qrcode"
	"github.com/nfnt/resize"
)

func main() {
	// URL or text for the QR code
	text := "https://ulbi.ac.id" // Replace with your ULBI campus URL

	// Generate QR code
	qrCode, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	// Open ULBI logo file
	logoFile, err := os.Open("logo_ulbi.png") // Replace with your ULBI logo file path
	if err != nil {
		fmt.Println("Error opening logo file:", err)
		return
	}
	defer logoFile.Close()

	// Decode ULBI logo
	logo, _, err := image.Decode(logoFile)
	if err != nil {
		fmt.Println("Error decoding logo image:", err)
		return
	}

	// Decode QR code image
	qrImage, _, err := image.Decode(bytes.NewReader(qrCode))
	if err != nil {
		fmt.Println("Error decoding QR code image:", err)
		return
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
	outFile, err := os.Create("qrcode.png") // Replace with desired output file name
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Encode the final image into the output file
	err = png.Encode(outFile, rgba)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	fmt.Println("QR code with ULBI logo generated successfully.")
}
