package fizzbuzz

type inMemCache struct {
  dict map[int]string
}

func NewInMemoryCache() Cache{
  return &inMemCache{
      dict: make(map[int]string),
  }
}

func (c *inMemCache) Put(key int, value string){
  c.dict[key] = value
}

func (c *inMemCache) Get(key int) (string, bool) {
  str, ok := c.dict[key]
  return str, ok
}
