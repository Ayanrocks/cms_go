package Database

import (
	"database/sql"

	"github.com/kataras/golog"

	_ "github.com/lib/pq"
)

func Connect() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		golog.Fatal(err)
	}

	defer func() {
		golog.Debug("Closing DB")
		db.Close()
	}()
}
