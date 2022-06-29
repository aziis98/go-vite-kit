package database

type memDB struct {}

func NewInMemoryDB() Database {
	return &memDB{}
}