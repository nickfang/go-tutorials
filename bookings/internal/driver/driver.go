package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDBLifetime = 5 * time.Minute

// ConnectSQl creates a connection to the database
func ConnectSQL(dsn string) (*DB, error) {
	fmt.Println("Connecting to database...", dsn)
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)
	dbConn.SQL = db

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB tries to ping the database
func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database connection
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
