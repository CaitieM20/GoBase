package fizzbuzz

//go:generate mockgen -source=cache.go -package=fizzbuzz -destination=cache_mock.go

type Cache interface {
  Put(key int, value string)
  Get(key int) (string, bool)
}