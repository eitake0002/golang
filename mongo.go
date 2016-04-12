package main

import(
  "fmt"
  "log"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct{
  ID   bson.ObjectId
  Name string 
  Age  int
}

func main(){
  fmt.Println("Start")
  log.Println("Start")

  session, _ := mgo.Dial("mongodb://localhost/humandb")
  defer session.Close()
  db := session.DB("humandb")
  fmt.Println(db)

  ritsu := &Person{
    Name: "田井中律",
    Age:  17,
  }
  col := db.C("humans")
  col.Insert(ritsu)

  p := new(Person)
  query := db.C("humans").Find(bson.M{})
  query.One(&p)

  fmt.Printf("%+v\n", p) 
}
