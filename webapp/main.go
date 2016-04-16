package main

import(
  "log"
  "net/http"
  "text/template"
  "./model/humans"
  //"gopkg.in/mgo.v2/bson"
)

type Page struct {
  Title string
  Counter int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  page := Page{"Hello World", 1}
  tmpl, err := template.ParseFiles("layout.html")
  if err != nil {
    panic(err)
  }
  err = tmpl.Execute(w, page)
  if err != nil {
    panic(err)
  }
}

type Human struct {
  //ID   bson.ObjectId
  Name string
  Age  int
}

func main(){
  //http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  //  w.Write([]byte(`<h1>test</h1>`))
  //})

  humans.Insert()
  log.Println("End")
  /*
  http.HandleFunc("/", viewHandler)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe", err)
  }
  */
}
