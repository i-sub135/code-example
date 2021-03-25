package main

import (
	"code-example/src/db"
	"code-example/src/store"
)

var (
	model    store.TProduk
	database db.RestDB
	sql      = database.Sql()
)

func main() {
	autoMigrate()
}

func autoMigrate() {
	sql.Migrator().DropTable(&model)
	sql.AutoMigrate(&model)
}
