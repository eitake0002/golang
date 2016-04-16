package humans

import (
  mgo "gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
)

type Human struct {
  Name string
  Age  int
}

func Insert() {
  session, _ := mgo.Dial("mongodb://localhost/humandb")
  defer session.Close()
  db := session.DB("humandb")

  human := Human{"test", 100}
  col := db.C("humans")
  col.Insert(human)
}
