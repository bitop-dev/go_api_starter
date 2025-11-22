package cache

// Cache abstracts a key/value store for optional Redis usage.
type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}
