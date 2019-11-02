package Database

import (
	"database/sql"

	"github.com/kataras/golog"

	_ "github.com/lib/pq"
)

func Connect() {
	connStr := "user=ayanrocks password=qwerty dbname=gocms sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		golog.Fatal(err)
	}

	row, err := db.Query(`INSERT INTO posts.posts(name, content) VALUES ('How to sample?', 'Post is vrry short')`)

	if err != nil {
		golog.Error(err)
	}

	golog.Debug(row)

	golog.Debug(db.Ping())

	defer func() {
		golog.Debug("Closing DB")
		db.Close()
	}()
}
