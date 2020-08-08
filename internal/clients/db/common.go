package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Parameters struct {
	Username string
	Password string
	Host     string
	Schema   string
}

type DB struct {
	Client *sql.DB
	Events []string
}

func New(params *Parameters) (*DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s search_path=%s sslmode=disable",
		params.Username, params.Password, params.Host, params.Schema)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}
