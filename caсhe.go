package golang_ninja

type MemoryCache[T any] struct {
	cacheMap map[string]T
}

func NewMemoryCache[T any]() *MemoryCache[T] {
	return &MemoryCache[T]{
		cacheMap: make(map[string]T),
	}
}

func (mc *MemoryCache[T]) Set(key string, value T) {
	mc.cacheMap[key] = value
}

func (mc *MemoryCache[T]) Get(key string) T {
	return mc.cacheMap[key]
}

func (mc *MemoryCache[T]) Del(key string) {
	delete(mc.cacheMap, key)
}
