package fizzbuzz

import (
  "github.com/golang/mock/gomock"
  "testing"
)


func Test_RunFizzBuzz_BadInput(t *testing.T) {
  mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()

  mockCache := NewMockCache(mockCtrl)
  handler := NewHandler(mockCache)

  _, err := handler.RunFizzBuzz([]string{"1", "2", "apple"})
  if err == nil {
    t.Error("Expected an error to be returned")
  }
}

func Test_RunFizzBuzz_CacheMiss(t *testing.T) {
  mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()

  mockCache := NewMockCache(mockCtrl)
  mockCache.EXPECT().Get(5).Return("", false)
  mockCache.EXPECT().Put(5, "Buzz")

  handler := NewHandler(mockCache)
  str, err := handler.RunFizzBuzz([]string{"5"})
  
  if err != nil {
    t.Error("Unexpected error returned", err)
  }
  if str[0] != "Buzz" {
    t.Error("Expected returned value to be Buzz", str)
  }
}

func Test_RunFizzBuzz_CacheHit(t *testing.T) {
  mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()

  mockCache := NewMockCache(mockCtrl)
  mockCache.EXPECT().Get(5).Return("Test", true)

  handler := NewHandler(mockCache)
  str, err := handler.RunFizzBuzz([]string{"5"})
  
  if err != nil {
    t.Error("Unexpected error returned", err)
  }

  if str[0] != "Test" {
    t.Error("Expected Test to be returned, not", str)
  }
}