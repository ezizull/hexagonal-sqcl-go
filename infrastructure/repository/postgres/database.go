package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
)

// Database cradential
var (
	// host     = os.Getenv("DB_HOST")
	// port     = os.Getenv("DB_PORT")
	// username = os.Getenv("DB_USERNAME")
	// password = os.Getenv("DB_PASSWORD")
	// dbname   = os.Getenv("DB_POSGRES")

	host     = "localhost"
	port     = 5432
	username = "root"
	password = "root"
	dbname   = "skyshi_gethired"
)

// NewSqlc is a function that returns a sqlc database connection
func NewSqlc() (*sqlc.Queries, error) {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, username, password, dbname)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	queries := sqlc.New(db)
	return queries, nil
}
