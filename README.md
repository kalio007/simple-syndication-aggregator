# Notes

## dev-dependencies

### sqcl 

for GO code from SQL
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### goose 
for migration
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

RUN
```
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="postgres://name:password@localhost:5432/db?sslmode=disable"
```

## DOCS

[SQCL](https://docs.sqlc.dev/)
