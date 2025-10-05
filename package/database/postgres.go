package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gilangnyan/point-of-sales/internal/config"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db   *sql.DB
	once sync.Once
}

func NewPostgresDB() Database {
	return &PostgresDB{}
}

func (p *PostgresDB) Connect(conf config.Config) error {
	var err error

	p.once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=%s",
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Name,
			conf.Database.SSLMode,
			conf.Database.TimeZone,
		)

		for i := 1; i <= 5; i++ {
			p.db, err = sql.Open("postgres", dsn)
			if err == nil && p.db.Ping() == nil {
				p.db.SetMaxOpenConns(25)
				p.db.SetMaxIdleConns(5)
				p.db.SetConnMaxLifetime(5 * time.Minute)
				p.db.SetConnMaxIdleTime(1 * time.Minute)

				log.Printf("Successfully connected to PostgreSQL database")
				return
			}

			if p.db != nil {
				p.db.Close()
				p.db = nil
			}

			log.Printf("Failed to connect to PostgreSQL database. Retrying in 5 seconds... (attempt %d/5)", i)
			time.Sleep(5 * time.Second)
		}

		err = errors.New("could not connect to PostgreSQL database after 5 attempts")
	})

	return err
}

func (p *PostgresDB) GetDB() *sql.DB {
	return p.db
}

func (p *PostgresDB) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

func (p *PostgresDB) Ping() error {
	if p.db != nil {
		return p.db.Ping()
	}
	return errors.New("database connection is nil")
}

func (p *PostgresDB) Health() (string, error) {
	if err := p.Ping(); err != nil {
		return "unhealthy", err
	}

	var version string
	err := p.db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return "connected but version check failed", err
	}

	return fmt.Sprintf("healthy - %s", version), nil
}

// Legacy function for backward compatibility (deprecated)
func Init(conf config.Config) (*sql.DB, error) {
	db := NewPostgresDB()
	err := db.Connect(conf)
	if err != nil {
		return nil, err
	}
	return db.GetDB(), nil
}
