package main

import "fmt"

func main () {
  array := [3]int{ 10, 20, 30 } // static sizing
  slice := []int{ 10, 20, 30, 40 } // dynamic sizing
  fmt.Println(array, slice)
  slice = append(slice, 50)
  // array = append(array, 50) // causes error
  fmt.Println(array, slice)
  fmt.Println(array[1])
  fmt.Println(slice[4])
  fmt.Println(array[1:])
  for i, x := range slice {
    fmt.Printf("slice[%d] = %d\n", i, x)
  }
}
