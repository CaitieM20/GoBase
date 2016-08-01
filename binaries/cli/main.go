package main

import (
  "fmt"
  "log"
  "github.com/CaitieM20/GoBase/fizzbuzz"
  "strconv"
  "os"
)

func main() {

  if len(os.Args) <= 1 {
    log.Fatal("Requires at least one integer as input")
  }

  nums, err := ConvertInput(os.Args[1:])

  if err != nil {
    log.Fatal("Input must be an integer", err)
  }

  fmt.Println(FizzBuzz(nums))
}

func ConvertInput(args []string) ([]int, error) {
  var rsp []int

  for _, str := range args {
    num, err := strconv.Atoi(str)

    if err != nil {
      return nil, err
    }

    rsp = append(rsp, num)
  }
  
  return rsp, nil 
}

func FizzBuzz(input []int) []string {
  var rsp []string 

  for _, num := range input {
    rsp = append(rsp, fizzbuzz.FizzBuzz(num))
  }

  return rsp
}
