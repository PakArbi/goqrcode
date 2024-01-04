package goqrcode 

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notifikasi struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Parkiran  	`json:"data"`
}

type Time struct {
	Message     string `json:"message,omitempty" bson:"message,omitempty"`
	WaktuMasuk  string `json:"waktumasuk,omitempty" bson:"waktumasuk,omitempty"`
	WaktuKeluar string `json:"waktukeluar,omitempty" bson:"waktukeluar,omitempty"`
}

type FormData struct {
	NamaLengkap    string `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	NPM            string `bson:"npm,omitempty" json:"npm,omitempty"`
	Jurusan        string `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	NamaKendaraan  string `bson:"namakendaraan,omitempty" json:"namakendaraan,omitempty"`
	NomorKendaraan string `bson:"nomorkendaraan,omitempty" json:"nomorkendaraan,omitempty"`
	JenisKendaraan string `bson:"jeniskendaraan,omitempty" json:"jeniskendaraan,omitempty"`
}

type Parkiran struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	ParkiranId     int                `json:"parkiranid" bson:"parkiranid"`
	Nama           string             `json:"nama" bson:"nama"`
	NPM            string             `json:"npm" bson:"npm"`
	Prodi        string             `json:"prodi" bson:"prodi"`
	NamaKendaraan  string             `json:"namakendaraan" bson:"namakendaraan"`
	NomorKendaraan string             `bson:"nomorkendaraan,omitempty" json:"nomorkendaraan,omitempty"`
	JenisKendaraan string             `json:"jeniskendaraan,omitempty" bson:"jeniskendaraan,omitempty"`
}

type Payload struct {
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

var stringnotif = []string{
	"Selamat Datang di PakArbi.silakan lakukan verifikasi di email ulbi Anda.",
}