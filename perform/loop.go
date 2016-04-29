package main 

import(
  "log"
  "time"
)

func main(){
  start := time.Now();
  for i := 0; i < 100000; i++ {
    
  }
  end := time.Now();
  log.Printf("%f sec\n",(end.Sub(start)).Seconds())
}
