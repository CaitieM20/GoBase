package fizzbuzz

import (
  "strconv"
)

type FizzBuzzHandler struct {
  cache Cache
}

func NewHandler(c Cache) *FizzBuzzHandler {
  return &FizzBuzzHandler{
    cache: c,
  }
}

func (h *FizzBuzzHandler) RunFizzBuzz(args[]string) ([]string, error) {

  nums, err := convertInput(args)
  if err != nil {
    return nil, err
  }

  var rsp []string 
  for _, num := range nums {
    // check the cache to see if we've already calculated it
    str, ok := h.cache.Get(num)
    
    // if not in the cache calculate it
    if !ok {
      str = fizzBuzz(num)
      h.cache.Put(num, str)
    }

    rsp = append(rsp, str)
  }

  return rsp, nil
}

func convertInput(args []string) ([]int, error) {
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
