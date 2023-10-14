package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func CreateDBConnection() (*sqlx.DB, *goqu.Database) {
	connectionStr := "postgres://postgres:2020@localhost:5432/krane?sslmode=disable"

	conn, err := sqlx.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	}
	return conn, goqu.New("postgres", conn)
}
