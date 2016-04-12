package main

import(
  "log"
  "math"
)

func Cosine_similarity(vector_list [][]int) []float64 {
  result_list := []float64{}
  for _, v1 := range vector_list {
    for _, v2 := range vector_list {
      sum := 0
      for int := 0; int < len(vector_list[0]); int++ {
        sum += v1[int] * v2[int]
      }
      result := math.Cbrt(float64(sum))
      result_list = append(result_list, result)
    }
  }
  return result_list
}

func Generate_vector(dic []string, sentence_words []string) [][]int {
  var vect_list [][]int
  sentence_vector := []int{}
  for _, dic_v := range dic {
    exists := 0
    for _, v := range sentence_words {
      if dic_v == v {
        exists = 1
      }
    }
    sentence_vector = append(sentence_vector, exists)
  }
  vect_list = append(vect_list, sentence_vector)
  return vect_list
}

func Create_dictionary(title_list [][]string) []string {
  dic := []string{}
  for _, words := range title_list {
    exists := 0
    for _, w := range words {
      log.Println(w)
      for _, d := range dic {
        if w == d {
          exists = 1
        }
      }
      if exists == 0 {
        dic = append(dic, w)
      }
      exists = 0
    }
  }
  return dic
}

func main(){
  title_list := [][]string{{"aaa","bbb"},{"aaa", "ccc"}}
  dic_result := Create_dictionary(title_list) 
  log.Println("dic_result :", dic_result)

  sample_input := [][]int{{1,0,1,0}, {1,1,1,1}}
  cos_list := Cosine_similarity(sample_input)
  log.Println(cos_list)

  dic := []string{"aaa", "bbb", "ccc"}
  sentence_words := []string{"bbb", "ccc"}
  result := Generate_vector(dic, sentence_words)
  log.Println(result)
}
