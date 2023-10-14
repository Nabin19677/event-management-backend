package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func CreateDBConnection() (*sql.DB, *goqu.Database) {
	connectionStr := "postgres://postgres:2020@localhost:5432/krane?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	}
	return conn, goqu.New("postgres", conn)
}
