package main

import (
  "fmt"
  "log"
  "github.com/CaitieM20/GoBase/fizzbuzz"
  "os"
)

func main() {

  if len(os.Args) <= 1 {
    log.Fatal("Requires at least one integer as input")
  }
  
  handler := fizzbuzz.NewHandler(
    &cache{
      dict: make(map[int]string),
  })

  nums, err := handler.RunFizzBuzz(os.Args[1:])
  if err != nil {
    log.Fatal("Input must be an integer", err)
  }

  fmt.Println(nums)
}

type cache struct {
  dict map[int]string
}

func (c *cache) Put(key int, value string){
  c.dict[key] = value
}

func (c *cache) Get(key int) (string, bool) {
  str, ok := c.dict[key]
  return str, ok
}

