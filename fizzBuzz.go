package main

import (
  "fmt"
)

func main() {
  for i := 1; i <= 100; i++ {
    fmt.Println(fizzbuzz(i))
  }
}

func isDivisibleBy(i int, num int) bool {
  return  i % num == 0
}

func fizzbuzz(i int) interface{} {
  if isDivisibleBy(i, 15) {
    return "FizzBuzz"
  } else if  isDivisibleBy(i, 3) {
    return "Fizz"
  } else if isDivisibleBy(i, 5) {
    return "Buzz"
  } else {
    return i
  }
}
