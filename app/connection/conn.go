package connection

import (
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

func NewConnection() *whatsapp.Conn {
	wac, err := whatsapp.NewConn(20 * time.Second)
	if err != nil {
		panic(err)
	}
	return wac
}
