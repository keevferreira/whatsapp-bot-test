package main

import (
	"github.com/eaceto/whatsapp-bot-test/app/connection"
)

func main() {
	whatsapp := connection.NewConnection()
	connection.Login(whatsapp)
}
