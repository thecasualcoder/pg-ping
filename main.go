package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

func main() {
	var conf PGConfig
	err := envconfig.Process("PGPING", &conf)
	if err != nil {
		log.Fatal(err)
	}

	if conf.Debug {
		json.NewEncoder(os.Stdout).Encode(conf)
	}

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
