package sqlitego

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Run() {
	db, err := sql.Open("sqlite", "./Sqlite/Go/bar.db")

	if err != nil {
		panic(err)
	}

	createTableSql := `
		CREATE TABLE IF NOT EXISTS foo (
			id integer not null primary key,
			name text
		);
	`

	res, err := db.Exec(createTableSql)

	if err != nil {
		panic(err)
	}

	fmt.Println(res.LastInsertId())

	insertSquerySql := `
		INSERT INTO foo (id, name) values (1, "André");
	`

	res, err = db.Exec(insertSquerySql)

	if err != nil {
		panic(err)
	}

	fmt.Println(res.LastInsertId())

	querySql := `
		SELECT * FROM foo WHERE id = ?;
	`

	type User struct {
		ID   int
		Nmae string
	}

	var u User

	err = db.QueryRow(querySql, 1).Scan(&u.ID, &u.Nmae)

	if err != nil {
		panic(err)
	}

	fmt.Println(u)
}
