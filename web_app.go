package main

import(
  "log"
  "net/http"
)

func main(){
  http.HandleFunc("/aaa", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`<h1>test</h1>`))
  })

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe", err)
  }
}
