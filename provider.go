package main

import "database/sql"

type SqliteProvider struct {
	db *sql.DB
}
