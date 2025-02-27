package port

type Cache interface {
	Set(key string, value []byte)
	Get(key string) ([]byte, bool)
	Delete(key string)
	Clear()
}
