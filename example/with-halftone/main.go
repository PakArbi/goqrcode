package main

import (
	"github.com/PakArbi/goqrcode/v2"
	"github.com/PakArbi/goqrcode/writer/standard"
)

func main() {
	qrc, err := qrcode.New("https://github.com/PakArbi/goqrcode")
	if err != nil {
		panic(err)
	}

	w0, err := standard.New("./repository_qrcode.png",
		standard.WithHalftone("./test.png"),
		standard.WithQRWidth(21),
	)
	handleErr(err)
	err = qrc.Save(w0)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
