package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type Storage struct {
	db *sql.DB
	mu sync.Mutex
}

const (
	fileName   = "storage/sso.db"
	driverName = "sqlite3"
)

func New() *Storage {
	conn := connectionMustLoad()
	return &Storage{db: conn}
}

func connectionMustLoad() *sql.DB {
	conn, err := sql.Open(driverName, fileName)
	if err != nil {
		panic("error while opening sqlite connection " + err.Error())
	}
	if err := initUsers(conn); err != nil {
		panic("error while initializing users table " + err.Error())
	}
	if err := initApps(conn); err != nil {
		panic("error while initializing apps table " + err.Error())
	}
	return conn
}

func initUsers(conn *sql.DB) error {
	const createQuery = `
      CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password_hash BLOB NOT NULL,
		is_admin BOOLEAN NOT NULL DEFAULT FALSE
	  );
	  CREATE INDEX IF NOT EXISTS idx_email ON users (email);`

	if _, err := conn.Exec(createQuery); err != nil {
		return err
	}
	return nil
}

func initApps(conn *sql.DB) error {
	const createQuery = `
      CREATE TABLE IF NOT EXISTS apps (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		secret TEXT
      );`

	if _, err := conn.Exec(createQuery); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
