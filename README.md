# css (citus-database-documentor)

This is a small utility I wrote while we were transitioning our Postgres
database to support sharding with [citus](https://github.com/citusdata/citus).
Since any tables you want to be distributed requires a sharding key, this
utility aims to at least begin documenting all of your tables and tries
to find the tenant_id/sharding key based on potential column names you provide.

## Install

```sh
$ go install github.com/whatl3y/cdd
```

## Parameters

1. `--connection` or `-c` [REQUIRED]: A postgres connection string to your DB we'll scan (i.e. postgres://localhost:5432/database)
2. `--tenant_id` or `-i` [REQUIRED]: A single or comma-delimited list of columns that will be your sharding key.
3. `--schema` or `-s` [OPTIONAL]: The schema we're looking in when scanning tables, DEFAULT: public

## Examples

```sh
# if '$GOPATH/bin' is in your path:
$ cdd -c postgres://localhost:5432/whatley?sslmode=disable -i tenantid
# Successfully created CSV to start using for your Citus migration: citus_info_############.csv

# if '$GOPATH/bin' is NOT in your path:
$ $GOPATH/bin/cdd -c postgres://localhost:5432/whatley?sslmode=disable -i tenantid
# Successfully created CSV to start using for your Citus migration: citus_info_############.csv

$ cdd --connection postgres://localhost:5432/whatley?sslmode=disable -tenant_id tenantid
# Successfully created CSV to start using for your Citus migration: citus_info_############.csv
```
