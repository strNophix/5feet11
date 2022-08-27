package db

import (
	"github.com/scylladb/gocqlx/v2"
)

func Seed(session gocqlx.Session) error {
	err := session.ExecStmt(`
		CREATE KEYSPACE IF NOT EXISTS fivefeeteleven 
		WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`)
	if err != nil {
		return err
	}

	err = session.ExecStmt(`
		CREATE TABLE IF NOT EXISTS fivefeeteleven.urls (
			id text PRIMARY KEY,
			long_url text
		)`)

	if err != nil {
		return err
	}

	return nil
}
