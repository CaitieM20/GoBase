package fizzbuzz

import (
  "github.com/leanovate/gopter"
  "github.com/leanovate/gopter/gen"
  "github.com/leanovate/gopter/prop"
  "testing"
)

func Test_ValidateFizzBuzz(t *testing.T) {
  parameters := gopter.DefaultTestParameters()
  parameters.MinSuccessfulTests = 10000
  properties := gopter.NewProperties(parameters)

  properties.Property("FizzBuzz Returns Correct String", prop.ForAll(
    func(num int) bool {

      str := fizzBuzz(num)

      switch str {
      case "Fizz":
        return (num % 3 == 0) && !(num % 5 == 0)
      case "Buzz":
        return (num % 5 == 0) && !(num % 3 == 0)
      case "FizzBuzz":
        return (num % 3 == 0) && (num % 5 == 0)
      default:
        return !(num % 3 == 0) && !(num % 5 == 0)
      }

    },
    gen.Int(),
  ))

  properties.TestingRun(t)
}