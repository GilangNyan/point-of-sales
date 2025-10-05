package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gilangnyan/point-of-sales/internal/config"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db   *sql.DB
	once sync.Once
}

func (m *MySQLDB) Connect(conf config.Config) error {
	var err error

	m.once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.Name,
			conf.Database.TimeZone,
		)

		for i := 1; i <= 5; i++ {
			m.db, err = sql.Open("mysql", dsn)
			if err == nil && m.db.Ping() == nil {
				m.db.SetMaxOpenConns(25)
				m.db.SetMaxIdleConns(5)
				m.db.SetConnMaxLifetime(5 * time.Minute)
				m.db.SetConnMaxIdleTime(1 * time.Minute)

				log.Printf("Successfully connected to MySQL database")
				return
			}

			if m.db != nil {
				m.db.Close()
				m.db = nil
			}

			log.Printf("Failed to connect to MySQL database. Retrying in 5 seconds... (attempt %d/5)", i)
			time.Sleep(5 * time.Second)
		}

		err = errors.New("could not connect to MySQL database after 5 attempts")
	})

	return err
}

func (m *MySQLDB) GetDB() *sql.DB {
	return m.db
}

func (m *MySQLDB) Close() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}

func (m *MySQLDB) Ping() error {
	if m.db != nil {
		return m.db.Ping()
	}
	return errors.New("database connection is nil")
}

func (m *MySQLDB) Health() (string, error) {
	if err := m.Ping(); err != nil {
		return "unhealthy", err
	}

	var version string
	err := m.db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "connected but version check failed", err
	}

	return fmt.Sprintf("healthy - MySQL %s", version), nil
}
