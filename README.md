# citus-database-documentor

This is a small utility I wrote while we were transitioning our Postgres
database to support sharding with [Citus](https://github.com/citusdata/citus).
Since for any tables you want to be distributed requires

## Install

```sh
$ go install github.com/whatl3y/citus-database-documentor
```

## Parameters

1. --connection or -c [REQUIRED]: A postgres connection string to your DB we'll scan (i.e. postgres://localhost:5432/database)
2. --tenant_id or -i [REQUIRED]: A single or comma-delimited list of columns that will be your sharding key.
3. --schema or -s [OPTIONAL]: The schema we're looking in when scanning tables, DEFAULT: public

## Examples

```sh
$ cdd -c postgres://localhost:5432/whatley?sslmode=disable -i site_id,site_id_int,idsite
# Successfully created CSV to start using for your Citus migration: citus_info_############.csv

$ cdd --connection postgres://localhost:5432/whatley?sslmode=disable -tenant_id site_id,site_id_int,idsite
# Successfully created CSV to start using for your Citus migration: citus_info_############.csv
```
