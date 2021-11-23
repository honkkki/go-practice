package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
	Id        int64          `db:"id"`
	Name      sql.NullString `db:"name"`
	Age       int            `db:"age"`
	CreatedAt string         `db:"created_at"`
}

func initDb() (err error) {
	dsn := "root:111111@tcp(127.0.0.1:3306)/golang"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(0)
	DB.SetMaxIdleConns(15) // 空闲连接数
	return nil
}

func queryRow() {
	var user User
	sqlStr := "select id, name from user where id = ?"
	row := DB.QueryRow(sqlStr, 1)
	// Scan内部已经defer close row
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}

func queryRows() {
	sqlStr := "select id, name from user"
	rows, _ := DB.Query(sqlStr)
	// ***必须关闭连接 否则会占用连接池的连接
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		fmt.Println(user)
	}

}

func main() {
	err := initDb()
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	queryRow()
	queryRows()
}
