package fizzbuzz

import (
  "testing"
)

func TestFizzBuzz_Fizz (t *testing.T) {

  if fizzBuzz(3) != "Fizz" {
    t.Error("Expected returned string to be 'Fizz'")
  }
}

func TestFizzBuzz_Buzz (t *testing.T) {
  if fizzBuzz(5) != "Buzz" {
    t.Error("Expected returned string to be 'Buzz'")
  }
}

func TestFizzBuzz_FizzBuzz (t *testing.T) {
  if fizzBuzz(15) != "FizzBuzz" {
    t.Error("Expected returned string to be 'FizzBuzz'")
  }
}

func TestFizzBuzz_Num (t *testing.T) {
  if fizzBuzz(2) != "2" {
    t.Error("Expected returned string to be '2'")
  }
}