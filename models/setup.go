package models

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var DB *bun.DB

func ConnectDatabase() {
	// dsn := "host=localhost user=postgres dbname=go_blog port=5432 sslmode=disable timezone=Asia/Kolkata"
    dsn := "postgres://postgres:@localhost:5432/go_blog?sslmode=disable"
    config, err := pgx.ParseConfig(dsn)
    if err != nil {
        panic(err)
    }
    config.PreferSimpleProtocol = true
    
    sqldb := stdlib.OpenDB(*config)
    DB = bun.NewDB(sqldb, pgdialect.New())
}