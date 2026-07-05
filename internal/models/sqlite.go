package models

import (
	"database/sql"

	moderncsqlite "modernc.org/sqlite"
)

func init() {
	for _, driverName := range sql.Drivers() {
		if driverName == "sqlite3" {
			return
		}
	}

	sql.Register("sqlite3", &moderncsqlite.Driver{})
}
