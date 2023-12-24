package connection

import (
	"fmt"

	whatsapp "github.com/Rhymen/go-whatsapp"
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
