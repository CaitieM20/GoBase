package main

import (
  "fmt"
  "log"
  "strconv"
  "os"
)

func main() {
    num, err := strconv.Atoi(os.Args[1])

    if err != nil {
      log.Fatal("Input must be an integer")
    }
    str := fizzBuzz(num)
    fmt.Println(str)
}

func fizzBuzz(num int) string {

  div3 := (num % 3) == 0
  div5 := (num % 5) == 0

  if div3 && div5 {
    return "FizzBuzz"
  } else if div3 {
    return "Fizz"
  } else if div5 {
    return "Buzz"
  } else {
    return strconv.Itoa(num)
  }
}