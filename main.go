package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

// PGConfig contains configurations to connect to a PG database
type PGConfig struct {
	Username      string
	Password      string
	Host          string
	DBName        string
	Query         string
	FrequencyInMS int32
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

func main() {
	var conf PGConfig
	err := envconfig.Process("PGPING", &conf)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(os.Stdout).Encode(conf)

	db, err := sql.Open("postgres", conf.ConnStr())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ticker := time.NewTicker(conf.GetFrequency())
	encoder := json.NewEncoder(os.Stdout)
	for range ticker.C {
		encoder.Encode(executeQuery(db, conf.GetQuery()))
	}
}

// SQLResult  tracks sql response and time taken
type SQLResult struct {
	Value          string
	TimeTakenInMS  float64
	Failed         bool
	FailureMessage string
}

func executeQuery(db *sql.DB, query string) SQLResult {
	res := SQLResult{}

	ctx, cancelFunc := context.WithCancel(context.Background())
	tenSecondTimer := time.NewTimer(10 * time.Second)
	go func() {
		<-tenSecondTimer.C
		cancelFunc()
	}()

	start := time.Now()
	rows, err := db.QueryContext(ctx, query)
	res.TimeTakenInMS = time.Since(start).Seconds() * 1000

	if err != nil {
		res.Failed = true
		res.FailureMessage = err.Error()
		return res
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			res.Failed = true
			res.FailureMessage = err.Error()
			return res
		}

		res.Value = name
	}

	return res
}
