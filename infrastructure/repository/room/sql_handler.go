package room

type SqlHandler interface {
	Find(obj interface{}, value ...interface{}) error
	First(obj interface{}, where ...interface{}) error
}
