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

func fizzbuzz(i int) interface{} {
  if i % 15 == 0 {
    return "FizzBuzz"
  } else if  i % 3 == 0 {
    return "Fizz"
  } else if i % 5 == 0{
    return "Buzz"
  } else {
    return i
  }
}
