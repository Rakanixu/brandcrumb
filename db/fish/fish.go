package fish

import (
	"github.com/Rakanixu/brandcrumb/models/fish"
)

/*
DB Interfaces for Fish

DB must be pluggable, allowing to change the implementation easily
(in memory, SQL like, non relational)

DBFish interface can be use whenever a implementation is imported too, then just use as a library.
*/

type DBFish interface {
	Creater
	Reader
	Deleter
	Searcher
}

type Creater interface {
	Create(t *fish.Fish) error
}

type Reader interface {
	Read(id int64) (*fish.Fish, error)
}

type Deleter interface {
	Delete(int64) error
}

type Searcher interface {
	Search(int64) ([]*fish.Fish, error)
}

var db DBFish

// Register implementation
func Register(dbFish DBFish) {
	db = dbFish
}

// Create fish
func Create(t *fish.Fish) error {
	return db.Create(t)
}

// Read fish
func Read(id int64) (*fish.Fish, error) {
	return db.Read(id)
}

// Delete fish
func Delete(id int64) error {
	return db.Delete(id)
}

// Search fish
func Search(tankId int64) ([]*fish.Fish, error) {
	return db.Search(tankId)
}
