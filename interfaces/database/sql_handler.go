package database

type SqlHandler interface {
	Find(obj interface{}, value ...interface{}) error
	First(obj interface{}, where ...interface{}) error
	Create(obj interface{}) error
	Save(obj interface{}) error
	Delete(obj interface{}, value ...interface{}) error
	Exec(sql string, value ...interface{}) error
	Raw(sql string, value ...interface{}) error
}
