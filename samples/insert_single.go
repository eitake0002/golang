package main

import(
  "log"
  "time"
  _ "fmt"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func main() {
  start := time.Now();
  Insert()
  end := time.Now();
  log.Printf("%f sec\n",(end.Sub(start)).Seconds())
}

func Insert() {
  log.Println("Start")
  db, err := sql.Open("mysql", "root:@/dev_test")
  defer db.Close()
  for i := 0; i < 10000; i++ {
    if err != nil {
      panic(err.Error())
    }

    qry := "insert into test_table (name) values ('this is test')"
    _, err := db.Exec(qry)
    if err != nil {
      panic(err.Error())
    }
    //log.Println(rows)
  }
}
