package main

import (
    //"fmt"
    "log"
    "database/sql"
    - "github.com/go-sql-driver/mysql"
    //"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
    //db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/app?charset=utf8")
    db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/app")
    db.SetMaxOpenConns(2000)
    db.SetMaxIdleConns(1000)
    db.Ping()
}

var (
    id int
    name string
)

func main() {
    rows, err := db.Query("SELECT uid, name FROM user limit 10")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        log.Println(id, name)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}
