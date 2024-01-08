package database

import (
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func returnDatabase() *sqlstore.Container {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("sqlite3", "file:tempDatabase.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	return container
}
