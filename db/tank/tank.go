package tank

import (
	"github.com/Rakanixu/brandcrumb/models/tank"
)

/*
DB Interfaces for Tank

DB must be pluggable, allowing to change the implementation easily
(in memory, SQL like, non relational)

DBTank interface can be use whenever a implementation is imported too, then just use as a library.
*/

type DBTank interface {
	Creater
	Reader
	Deleter
}

type Creater interface {
	Create(t *tank.Tank) error
}

type Reader interface {
	Read(id int64) (*tank.Tank, error)
}

type Deleter interface {
	Delete(int64) error
}

var db DBTank

// Register implementation
func Register(dbTank DBTank) {
	db = dbTank
}

// Create tank
func Create(t *tank.Tank) error {
	return db.Create(t)
}

// Read tank
func Read(id int64) (*tank.Tank, error) {
	return db.Read(id)
}

// Delete tank
func Delete(id int64) error {
	return db.Delete(id)
}
