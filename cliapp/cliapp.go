package cliapp

import (
	"errors"
	"sort"

	"github.com/urfave/cli"
)

type Args struct {
	Connection string
	Schema     string
	TenantID   string
}

func CreateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "citus-database-documentor"
	app.Usage = "generate a CSV of information useful when converting your DB schema to support sharding."
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Lance Whatley",
			Email: "whatl3y@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "connection, c",
			Usage:  "connection string to the postgres DB you are connecting to",
			EnvVar: "DATABASE_URL",
		},
		cli.StringFlag{
			Name:  "schema, s",
			Usage: "Schema for the database that we are getting tables from.",
		},
		cli.StringFlag{
			Name:  "tenant_id, i",
			Usage: "Either a single or comma-delimited list of columns that represent the tenant ID for sharding.",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}

func ValidateInputs(c *cli.Context) (Args, error) {
	connection := c.String("connection")
	schema := c.String("schema")
	tenantID := c.String("tenant_id")

	if connection == "" {
		return Args{"", "", ""}, errors.New("Please provide a connection string (-c or --connection)")
	}

	if tenantID == "" {
		return Args{"", "", ""}, errors.New("Please provide the column or columns that represent the tenant ID for sharding (-i or --tenant_id)")
	}

	return Args{connection, schema, tenantID}, nil
}
