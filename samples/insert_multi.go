package main

import(
  "log"
  "time"
  _ "fmt"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func main() {
  log.Println("Start")
  start := time.Now();
  pal_num := 100
  ch := make(chan int, pal_num)
  for i := 0; i < pal_num; i++ {
    go Insert(ch)
  }
  for i := 0; i < pal_num; i++ {
    <-ch
  }
  end := time.Now();
  log.Printf("%f sec\n",(end.Sub(start)).Seconds())
}

func Insert(ch chan int) {
  db, err := sql.Open("mysql", "root:@/dev_test")
  defer db.Close()
  for i := 0; i < 100; i++ {
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
  ch <- 1
}
