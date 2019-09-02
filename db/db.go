/*
 *
 */
package db

var (
	dbMap = map[string]*Cache{}
)

func NewDB(name string) {
	_, ok := dbMap[name]
	if !ok {
		dbMap[name] = New(1024, LFU)
	}
}

func GetDB(name string) *Cache {
	return dbMap[name]
}
