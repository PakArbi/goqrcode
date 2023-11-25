package goqrcode 

type Notif struct {
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

type FormData struct {
	NamaLengkap    string `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	NPM            string `bson:"npm,omitempty" json:"npm,omitempty"`
	Jurusan        string `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	NamaKendaraan  string `bson:"namakendaraan,omitempty" json:"namakendaraan,omitempty"`
	NomorKendaraan string `bson:"nomorkendaraan,omitempty" json:"nomorkendaraan,omitempty"`
	JenisKendaraan string `bson:"jeniskendaraan,omitempty" json:"jeniskendaraan,omitempty"`
}

type simpleQRCode struct {
    Content 	string	`bson:"content,omitempty" json:"content,omitempty"`
    Size        int		`bson:"size,omitempty" json:"size,omitempty"`
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