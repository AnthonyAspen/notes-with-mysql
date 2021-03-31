package main

import (
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "fmt"
)
const (
  db_user = "root"
  db_passwd = "root"
  db_addr = "localhost"
  db_db = "Notes"
)
func getConnectionToMYSQL() (db *sql.DB,err error){
  s :=  fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",db_user,db_passwd,db_addr,db_db)
  db,err = sql.Open("mysql",s)
  return db,err

}
