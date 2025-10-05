package database

import (
	"database/sql"
	"fmt"
	"gilangnyan/point-of-sales/internal/config"
)

// Database interface defines the contract for database operations
type Database interface {
	// Connect establishes connection to the database
	Connect(config config.Config) error

	// GetDB returns the underlying database connection
	GetDB() *sql.DB

	// Close closes the database connection
	Close() error

	// Ping checks if the database connection is alive
	Ping() error

	// Health returns database health status
	Health() (string, error)
}

// DatabaseType represents different database types
type DatabaseType string

const (
	PostgreSQL DatabaseType = "postgresql"
	MySQL      DatabaseType = "mysql"
	SQLite     DatabaseType = "sqlite"
)

// NewDatabase creates a new database instance based on the specified type
func NewDatabase(dbType DatabaseType) (Database, error) {
	switch dbType {
	case PostgreSQL:
		return &PostgresDB{}, nil
	case MySQL:
		return &MySQLDB{}, nil
	case SQLite:
		return nil, fmt.Errorf("SQLite implementation not yet available")
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

// ConnectDatabase is a helper function to create and connect database in one step
func ConnectDatabase(dbType DatabaseType, conf config.Config) (Database, error) {
	db, err := NewDatabase(dbType)
	if err != nil {
		return nil, err
	}

	if err := db.Connect(conf); err != nil {
		return nil, err
	}

	return db, nil
}
