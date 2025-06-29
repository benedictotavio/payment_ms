package db

import (
	"database/sql"
	"strings"
)

type ConfigPsql struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Psql struct {
	DB *sql.DB
}

func NewDB(config ConfigPsql) (*Psql, error) {
	db, err := connectDatabase(config)
	if err != nil {
		return nil, err
	}
	return &Psql{DB: db}, nil
}

func (p *Psql) Close() error {
	return p.DB.Close()
}

func connectDatabase(config ConfigPsql) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		buildDatabaseString(config),
	)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	pingErr := db.Ping()

	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}

func buildDatabaseString(config ConfigPsql) string {
	var sb strings.Builder
	sb.WriteString("host=")
	sb.WriteString(config.Host)
	sb.WriteString(" port=")
	sb.WriteString(config.Port)
	sb.WriteString(" user=")
	sb.WriteString(config.User)
	sb.WriteString(" password=")
	sb.WriteString(config.Password)
	sb.WriteString(" dbname=")
	sb.WriteString(config.DBName)
	sb.WriteString(" sslmode=disable")
	return sb.String()
}