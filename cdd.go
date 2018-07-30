package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/urfave/cli"
	"github.com/whatl3y/citus-database-documentor/cliapp"
)

func main() {
	csvFileName := "citus_info_" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".csv"
	app := cliapp.CreateApp()

	app.Action = func(c *cli.Context) error {
		args, err := cliapp.ValidateInputs(c)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		var possibleTenantIds = strings.Split(strings.Replace(args.TenantID, " ", "", -1), ",")
		db, err := sql.Open("postgres", args.Connection)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		schema := "public"
		if args.Schema != "" {
			schema = args.Schema
		}

		rows, err := db.Query(`
			SELECT table_name, column_name, data_type
			FROM information_schema.columns
			WHERE table_schema = $1
			ORDER BY table_name
		`, schema)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		defer rows.Close()

		citusCsvData := [][]string{{"table_name", "tenant_id", "tenant_id_data_type", "distributed_query", "reference_query"}}

		isFirstTable := true
		var currentTable, tenantIDColumn, currentTenantIDDataType string = "", "N/A", "N/A"
		for rows.Next() {
			var tableName, columnName, dataType string
			if err := rows.Scan(&tableName, &columnName, &dataType); err != nil {
				log.Fatal(err)
			}

			if currentTable != tableName {
				if isFirstTable == true {
					isFirstTable = false
				} else {
					var distQuery, refQuery string = "SELECT create_distributed_table('" + currentTable + "', '" + tenantIDColumn + "')", "SELECT create_reference_table('" + currentTable + "')"
					if tenantIDColumn == "N/A" {
						distQuery = ""
						refQuery = ""
					}
					citusCsvData = append(citusCsvData, []string{currentTable, tenantIDColumn, currentTenantIDDataType, distQuery, refQuery})
				}
				currentTable = tableName
				tenantIDColumn = "N/A"
				currentTenantIDDataType = "N/A"
			} else if isPossibleTenantId(possibleTenantIds, columnName) {
				tenantIDColumn = columnName
				currentTenantIDDataType = dataType
			}
		}

		fileWriter, err := os.Create(csvFileName)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		w := csv.NewWriter(fileWriter)
		for _, record := range citusCsvData {
			if err := w.Write(record); err != nil {
				log.Fatalln("Error writing entry to csv:", err)
			}
		}

		// Write any buffered data to the underlying writer (standard output).
		w.Flush()
		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Successfully created CSV to start using for your Citus migration:", csvFileName)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func isPossibleTenantId(tenantIds []string, possibility string) bool {
	for _, str := range tenantIds {
		if str == possibility {
			return true
		}
	}
	return false
}
