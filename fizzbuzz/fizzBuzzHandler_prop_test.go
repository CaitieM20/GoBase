package fizzbuzz

import (
  "strconv"
  "github.com/leanovate/gopter"
  "github.com/leanovate/gopter/gen"
  "github.com/leanovate/gopter/prop"
  "testing"
)

func Test_ValidateRunFizzBuzz(t *testing.T) {
  parameters := gopter.DefaultTestParameters()
  parameters.MinSuccessfulTests = 100
  properties := gopter.NewProperties(parameters)

  properties.Property("FizzBuzz Returns Correct String", prop.ForAll(
    func(requests [][]string) bool {
      handler := NewHandler(NewInMemoryCache())

      for _, nums := range requests {
        results, _ := handler.RunFizzBuzz(nums)

        for i, str := range results {
          num, _ := strconv.Atoi(nums[i])

          switch str {
          case "Fizz":
            return (num % 3 == 0) && !(num % 5 == 0)
          case "Buzz":
            return (num % 5 == 0) && !(num % 3 == 0)
          case "FizzBuzz":
            return (num % 3 == 0) && (num % 5 == 0)
          default:
            expectedStr := strconv.Itoa(num)
            return !(num % 3 == 0) && !(num % 5 == 0) && expectedStr == str
          }
        }
      }

      return true
    },
    gen.SliceOf(gen.SliceOf(gen.NumString())),
  ))

  properties.TestingRun(t)
}