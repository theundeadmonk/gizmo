package postgresql

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/NYTimes/gizmo/config"
)

// Config holds everything yo need to
// connect and interact with a Postgresql DB
type Config struct {
	User    string `envconfig:"POSTGRESQL_USER"`
	Pw      string `envconfig:"POSTGRESQL_PW"`
	Host    string `envconfig:"POSTGRESQL_HOST_NAME"`
	Port    int    `envconfig:"POSTGRESQL_PORT"`
	DBName  string `envconfig:"POSTGRESQL_DB_NAME"`
	SSLMode string `envconfig:"POSTGRESQL_SSL_MODE"`
}

const (
	// DefaultSSLMode is disabled
	DefaultSSLMode = "disable"
	// DefaultPort is the default post for Postgresql connections
	DefaultPort = 5432
)

// DB will open a sql connection.
// Users must import a postgresql driver in their
// main to use this.
func (p *Config) DB() (*sql.DB, error) {
	db, err := sql.Open("postgres", p.String())
	if err != nil {
		return db, err
	}
	return db, nil
}

// String will return the Postgresql connection string
func (m *Config) String() string {
	if m.Port == 0 {
		m.Port = DefaultPort
	}

	if m.SSLMode != "" {
		m.SSLMode = url.QueryEscape(m.SSLMode)
	} else {
		m.SSLMode = url.QueryEscape(DefaultSSLMode)
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%i/%s?sslmode=%s",
		m.User,
		m.Pw,
		m.Host,
		m.Port,
		m.DBName,
		m.SSLMode,
	)
}

// LoadConfigFromEnv will attempt to load a Postgresql object
// from environment variables. If not populated, nil
// is returned
func LoadCOnfigFromEnv() *Config {
	var postgres Config
	config.LoadEnvConfig(&postgres)
	if postgres.Host != "" {
		return &postgres
	}
	return nil
}
