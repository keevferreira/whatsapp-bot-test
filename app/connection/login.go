package connection

import (
	"fmt"

	whatsapp "go.mau.fi/whatsmeow"
)

func Login(whatsapp *whatsapp.Conn) whatsapp.Session {
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
	}()
	session, err := whatsapp.Login(qrChan)
	if err != nil {
		panic(err)
	}
	return session
}
