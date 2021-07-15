package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var sqlxDB *sqlx.DB

type UserInfo struct {
	Id        int64  `db:"id"`
	Name      sql.NullString `db:"name"`
	Age       int    `db:"age"`
	CreatedAt sql.NullString `db:"created_at"`
}

func initSqlxDb() (err error) {
	dsn := "root:111111@tcp(127.0.0.1:3306)/golang"
	sqlxDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	//sqlxDB.SetMaxOpenConns(100)
	sqlxDB.SetMaxIdleConns(15)
	return nil
}

func queryData() {
	var user UserInfo
	sqlStr := "select name from user where id = ?"
	err := sqlxDB.Get(&user, sqlStr, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Name.String)
}

func main() {
	err := initSqlxDb()
	if err != nil {
		log.Fatal(err)
	}

	queryData()
}
