package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/thecasualcoder/pg-ping/pkg/pg"
	"github.com/urfave/cli"
)

// Execute the app
func Execute(app *cli.App) error {
	app.Name = "pg-ping"
	app.Usage = "Ping your postgres continously"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "once",
			Usage:  "Ping only once and quit",
			EnvVar: "PGPING_ONCE",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Print debug logs",
			EnvVar: "PGPING_DEBUG",
		},
		cli.StringFlag{
			Name:   "username, U",
			Usage:  "Username to connect to postgres",
			EnvVar: "PGPING_USERNAME",
		},
		cli.StringFlag{
			Name:   "password, p",
			Usage:  "Password to connect to postgres",
			EnvVar: "PGPING_PASSWORD",
		},
		cli.StringFlag{
			Name:   "host, h",
			Usage:  "Host of postgres server",
			EnvVar: "PGPING_HOST",
		},
		cli.StringFlag{
			Name:   "dbname, d",
			Usage:  "DBName to connect to",
			EnvVar: "PGPING_DBNAME",
		},
		cli.IntFlag{
			Name:   "frequency, f",
			Usage:  "Frequency to ping",
			EnvVar: "PGPING_FREQUENCY_IN_MS",
		},
		cli.StringFlag{
			Name:   "query, c",
			Usage:  "Query to run",
			EnvVar: "PGPING_QUERY",
			Value:  "SELECT 1",
		},
	}
	return app.Run(os.Args)
}

func run(c *cli.Context) error {
	if len(c.Args()) > 0 {
		cli.ShowAppHelp(c)
		return fmt.Errorf("args are not allowed")
	}
	conf := pg.Config{
		Username:      c.String("username"),
		Password:      c.String("password"),
		Host:          c.String("host"),
		DBName:        c.String("dbname"),
		FrequencyInMS: int32(c.Int("frequency")),
		Query:         c.String("query"),
	}

	encoder := json.NewEncoder(os.Stdout)

	if c.BoolT("debug") {
		encoder.Encode(conf)
	}

	db, err := pg.NewDB(conf)
	if err != nil {
		return err
	}
	defer db.Close()

	var resultChan <-chan pg.SQLResult

	if c.BoolT("once") {
		resultChan = db.PingOnce()
	} else {
		resultChan = db.Ping()
	}

	formatter := prettyjson.NewFormatter()

	formatter.Newline = ""
	formatter.Indent = 0

	for r := range resultChan {
		data, _ := formatter.Marshal(r)
		fmt.Println(string(data))
	}

	return nil
}
