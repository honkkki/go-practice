package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var sqlxDB *sqlx.DB

type UserInfo struct {
	ID        int64          `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	Age       int            `db:"age" json:"age"`
	CreatedAt sql.NullString `db:"created_at" json:"created_at"`
}

func initSqlxDb() (err error) {
	dsn := "root:111111@tcp(127.0.0.1:3306)/golang"
	sqlxDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	sqlxDB.SetMaxOpenConns(100)
	sqlxDB.SetMaxIdleConns(15)
	return nil
}

func queryData() {
	var user UserInfo
	sqlStr := "select * from user where id = ?"
	err := sqlxDB.Get(&user, sqlStr, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", user)
	//fmt.Println(user.CreatedAt.String)

	js, _ := json.Marshal(struct {
		ID   int64  `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
		Age  int    `db:"age" json:"age"`
	}{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	})
	fmt.Println(string(js))
}

func selectData() {
	var data []UserInfo
	sqls := "select * from user"
	err := sqlxDB.Select(&data, sqls)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", data)
}

func save() {
	user := UserInfo{
		Name: "karina",
		Age:  18,
		CreatedAt: sql.NullString{
			String: time.Now().Format("2006-01-02 15:04:05"),
			Valid:  true,
		},
	}

	insertSql := "insert into user (id,name,age,created_at) values (:id, :name, :age, :created_at)"

	_, err := sqlxDB.NamedExec(insertSql, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success")
}

func main() {
	err := initSqlxDb()
	if err != nil {
		log.Fatal(err)
	}

	queryData()
	//save()
	selectData()
}
