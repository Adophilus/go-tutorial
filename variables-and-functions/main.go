package main

import "fmt"

func formatGreeting (username string) string {
  return fmt.Sprintf("Welcome %s", username)
}

func greetUser (username string) {
  fmt.Println(formatGreeting(username))
}

func addN (n int) func (x int) int {
  return func (x int) int {
    return n + x
  }
}

func isEven (n int) (bool, error) {
  if n <= 0 {
    return false, fmt.Errorf("error: n must be > 0")
  }
  return n % 2 == 0, nil
}

func main () {
  var x int = 10
  name := "user"
  fmt.Println("The value of x is", x)
  greetUser(name)
  fmt.Println(addN(2)(4))
  is_even, err := isEven(x)
  if err == nil {
    fmt.Println(fmt.Sprintf("%d is divisible by 2?", x), is_even)
  }
}
