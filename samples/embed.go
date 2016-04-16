package main

import(
  "log"
)

type WrapStruct struct {
  MyStruct
  wrapper_name string
}

type MyStruct struct {
  name string
  age int
}

func main(){
  my_struct := MyStruct{"test", 20}
  wrapper := WrapStruct{my_struct, "wrapper string value"}
  log.Println(wrapper.MyStruct.name)
  log.Println(wrapper.wrapper_name)
}
