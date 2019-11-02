package Database

import (
	"database/sql"

	"github.com/kataras/golog"

	_ "github.com/lib/pq"
)

// DB Variable
var DB *sql.DB

// Connect Function
func Connect() {
	connStr := "user=ayanrocks password=qwerty dbname=gocms sslmode=disable"

	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		golog.Error(err)
	}

	golog.Debug("DataBase Connected ", DB)
	// golog.Debug(DB.Ping())

	// defer func() {
	// 	golog.Debug("Closing DB")
	// 	DB.Close()
	// }()

	defer DB.Close()
}
