package fizzbuzz

import (
  "testing"
)

func TestFizzBuzz_Fizz (t *testing.T) {

  if FizzBuzz(3) != "Fizz" {
    t.Error("Expected returned string to be 'Fizz'")
  }
}

func TestFizzBuzz_Buzz (t *testing.T) {
  if FizzBuzz(5) != "Buzz" {
    t.Error("Expected returned string to be 'Buzz'")
  }
}

func TestFizzBuzz_FizzBuzz (t *testing.T) {
  if FizzBuzz(15) != "FizzBuzz" {
    t.Error("Expected returned string to be 'FizzBuzz'")
  }
}

func TestFizzBuzz_Num (t *testing.T) {
  if FizzBuzz(2) != "2" {
    t.Error("Expected returned string to be '2'")
  }
}