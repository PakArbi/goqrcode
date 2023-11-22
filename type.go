package goqrcode 

type Notif struct {
	Message string `bson:"message"`
}

// QRCodeVerification adalah tipe data untuk verifikasi kode QR
type QRCodeVerification struct {
	OriginalData string // Data asli yang akan divalidasi dengan kode QR
	QRData       string // Data dari kode QR yang akan divalidasi
	IsValid      bool   // Status validasi kode QR
}

type EmailData struct {
	Email string `json:"email"`
}

type Payload struct {
	Email string `json:"email"`
	Message string `json:"message"`
}

type Database interface {
	InsertPayload(payload Payload) error
}

type MockDatabase struct {
	InsertPayloadFunc func(payload Payload) error
}

var stringnotif = []string{
	"Selamat Datang di PakArbi.silakan lakukan verifikasi di email ulbi Anda.",
}