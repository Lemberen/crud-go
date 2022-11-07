package dao

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"

	"crud/model/db/queries"
)

var db, err = sql.Open("sqlite", "./model/db/bookstore.db")

func init() {
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	migrateDb()
}

func Delete(id int) sql.Result {
	result, err := db.Exec(queries.Delete, id)

	if err != nil {
		log.Fatal("DELETE FAILED: ", err)
	}

	return result
}

func Get() *sql.Rows {
	result, err := db.Query(queries.Get)

	if err != nil {
		log.Fatal("GET FAILED:", err)
	}

	return result
}

func Post(id int, title string, author string, publisher string, language string, pages int, edition int, year int, isbn string) sql.Result {
	result, err := db.Exec(queries.Post, id, title, author, publisher, language, pages, edition, year, isbn)

	if err != nil {
		log.Fatal("INSERT FAILED: ", err)
	}

	return result
}

func Put(id int, title string, author string, publisher string, language string, pages int, edition int, year int, isbn string) sql.Result {
	result, err := db.Exec(queries.Put, id, title, author, publisher, language, pages, edition, year, isbn)

	if err != nil {
		log.Fatal("UPDATE FAILED: ", err)
	}

	return result
}

func migrateDb() {
	m, err := migrate.New("file://./model/db/migrations", "sqlite3://./model/db/bookstore.db")

	if err != nil {
		log.Fatal("MIGRATION FAILED: ", err)
	}

	m.Up()
}
