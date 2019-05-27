package models

import (
	db "databases"

)


func GetPersonList(name string) db.Book {
	book := db.Find(name)

	return  book
}