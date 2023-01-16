package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var db *bun.DB

func Connect() {
	sqldb, err := sql.Open(sqliteshim.ShimName, "postgres://localhost:5432/blog-master")
	if err != nil {
		panic(err)
	}
	db = bun.NewDB(sqldb, sqlitedialect.New())
}

func Get() *bun.DB {
	return db
}
