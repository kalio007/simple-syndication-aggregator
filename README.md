Notes
dev-dependencies
### sqcl 

for GO code fro SQl
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### goose 
 SQL for migration
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="postgres://name:password@localhost:5432/db?sslmode=disable"
