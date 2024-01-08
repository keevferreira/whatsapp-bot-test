package main

import (
	"github.com/keevferreira/whatsapp-bot-test/app/connection"
)

func main() {
	whatsapp := connection.NewConnection()
	connection.Login(whatsapp)
}
