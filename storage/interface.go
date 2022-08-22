package storage

type Storer interface {
	Get(key string) bool
	Put(key string, value bool)
	Delete(key string)
	GetAndDelKey() string
}
