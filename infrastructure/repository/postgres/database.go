package postgres

import (
	"hexagonal-sqlc/infrastructure/repository/postgres/sqlc"

	_ "github.com/lib/pq"
)

// NewSqlc is a function that returns a sqlc database connection
func NewSqlc() (*sqlc.Queries, error) {
	var infoPg infoDatabasPostgres
	err := infoPg.getPostgresConn("Databases.PostgreSQL.Localhost")
	if err != nil {
		return nil, err
	}

	sqlDB, err := initPostgresDB(infoPg)
	if err != nil {
		return nil, err
	}

	queries := sqlc.New(sqlDB)
	return queries, nil
}
