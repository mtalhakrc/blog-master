package database

import (
	"database/sql"
	"fmt"
	"github.com/blog-master/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var db *bun.DB

func Connect(cfg config.DatabaseConfig) {
	sqldb, err := sql.Open(sqliteshim.ShimName, fmt.Sprintf("posgres://%s:%s/%s", cfg.Host, cfg.Port, cfg.Name))
	if err != nil {
		panic(err)
	}
	db = bun.NewDB(sqldb, sqlitedialect.New())
}

func Get() *bun.DB {
	return db
}
