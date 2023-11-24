package goqrcode 

type Notif struct {
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

// QRCodeVerification adalah tipe data untuk verifikasi kode QR
type QRCodeVerification struct {
	OriginalData string // Data asli yang akan divalidasi dengan kode QR
	QRData       string // Data dari kode QR yang akan divalidasi
	IsValid      bool   // Status validasi kode QR
}

type EmailData struct {
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}

type Payload struct {
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

type Database interface {
	InsertPayload(payload Payload) error
}

// type MockDatabase struct {
// 	InsertPayloadFunc func(payload Payload) error
// }

type MockDB struct {
	EmailSent bool
}

// // MockEmailSender struct to mock email sending functionality
// type MockEmailSender struct {
// 	VerificationEmailSent bool
// }

var stringnotif = []string{
	"Selamat Datang di PakArbi.silakan lakukan verifikasi di email ulbi Anda.",
}