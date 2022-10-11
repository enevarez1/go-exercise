package db

import "database/sql"

// Provides all functions to use db queries
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}