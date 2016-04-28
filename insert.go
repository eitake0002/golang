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
  //ch := make(chan int, 10)
  for i := 0; i < 10; i++ {
    Insert(ch)
  }
  //for i := 0; i < 10; i++ {
  //  <-ch
  //}
  end := time.Now();
  log.Printf("%f sec\n",(end.Sub(start)).Seconds())
}

func Insert() {
  log.Println("Start")
  db, err := sql.Open("mysql", "root:@/dev_test")
  defer db.Close()
  for i := 0; i < 1000; i++ {
    if err != nil {
      panic(err.Error())
    }

    qry := "insert into test_table (name) values ('this is test')"
    rows, err := db.Exec(qry)
    if err != nil {
      panic(err.Error())
    }
    log.Println(rows)
  }
  //ch <- 1
}
