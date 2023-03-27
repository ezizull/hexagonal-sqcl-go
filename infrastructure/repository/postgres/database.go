package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
)

// DB cradential
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

func New() *sqlc.Queries {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, username, password, dbname)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to PostgresDB:", err)
	}

	queries := sqlc.New(db)
	return queries
}
