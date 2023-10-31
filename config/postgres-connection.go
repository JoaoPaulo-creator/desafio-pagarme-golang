package config

import (
	"database/sql"

	"github.com/desafio-pagarme-golang/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

type DatabaseConnection struct {
	DB *sql.DB
	Q  *db.Queries
}

func (d *DatabaseConnection) Close() error {
	return d.DB.Close()
}

func NewPostgresConnection() (*DatabaseConnection, error) {
	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:banco123@localhost:5432/pagarme?sslmode=disable",
	)
	if err != nil {
		return nil, err
	}

	db := db.New(conn)
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://sql/migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return &DatabaseConnection{conn, db}, nil
}
