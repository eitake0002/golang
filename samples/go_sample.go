package main

import(
  "log"
)

type Samples struct {
  Id      int
  Name    string
}

func (s *Samples) SliceSample() {
  sample_slice := make([]int, 3)
  for i, v := range sample_slice {
    log.Println(i, v)
  }
}

func (s *Samples) MapSample() {
  sample_map := map[int] string{1: "val1", 2: "val2"}
  for i, v := range sample_map {
    log.Println(i, v)
  }
}

func main(){
  s := Samples{Id: 1, Name: "sample_name"}
  s.SliceSample()
  s.MapSample()
}
