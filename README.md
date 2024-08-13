# go-rest-api
A Resful API in go using GIN and GORM

### Environment Variables

Create a .env in the project root folder
```
PORT=3000
DB_URL=db.sqlite3
```

Open terminal and change to project root folder
```shell
go mod download
go run migrate/migrate.go
go run main.go
```
