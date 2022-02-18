package main

import "fmt"

func main () {
  count := map[string]int{
    "the": 400,
    "and": 600,
    "maybe": 10,
  }
  count["therefore"] = 30
  fmt.Println(count["the"])
  for key, value := range count {
    fmt.Printf("%s occurred %d times\n", key, value)
  }

  distrib := make(map[int]int) // make() is critical when starting with an empty map
  distrib[5] = 1e9
  distrib[10] = 1e8
  distrib[15] = 1e7
  distrib[20] = 1e5
  for age, freq := range distrib {
    fmt.Printf("%d year olds are %d in number\n", age, freq)
  }
}
