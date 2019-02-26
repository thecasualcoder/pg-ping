package main

import (
	"fmt"
	"time"
)

// PGConfig contains configurations to connect to a PG database
type PGConfig struct {
	Username      string
	Password      string
	Host          string
	DBName        string
	Query         string
	FrequencyInMS int32
	Debug         bool
}

// ConnStr returns a connection string to connect to postgres
func (c *PGConfig) ConnStr() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Username, c.Password, c.Host, c.DBName)
}

// GetQuery returns the query to use to ping
func (c *PGConfig) GetQuery() string {
	if c.Query == "" {
		return "select 1"
	}

	return c.Query
}

// GetFrequency returns the frequence in MS in which the query should be run
func (c *PGConfig) GetFrequency() time.Duration {
	if c.FrequencyInMS == 0 {
		return time.Second
	}

	return time.Duration(c.FrequencyInMS) * time.Millisecond
}
