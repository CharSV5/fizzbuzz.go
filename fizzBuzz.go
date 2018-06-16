package main

// import (
//   "fmt"
// )

// func main() {
//   isDivisibleBy(30, 15)
// }

func isDivisibleBy(i int, num int) bool {
  return  i % num == 0
}

func isDivisibleBy15(i int) string {
  if i % 15 == 0 {
    return "FizzBuzz"
  } else {
    return ""
  }
}
