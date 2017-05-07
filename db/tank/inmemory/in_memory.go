package inmemory

import (
	"errors"
	dbFish "github.com/Rakanixu/brandcrumb/db/fish"
	dbTank "github.com/Rakanixu/brandcrumb/db/tank"
	"github.com/Rakanixu/brandcrumb/models/tank"
	"net/http"
	"strconv"
	"sync"
)

// memory holds all tanks in the system
// it is an approach to a relational DB, so this represents the tank table
type memory struct {
	mutex   sync.Mutex
	records map[int64]*tank.Tank
	count   int64
}

// init auto register implementation when imported
func init() {
	dbTank.Register(&memory{
		records: make(map[int64]*tank.Tank),
	})
}

// Create tank in memory
func (m *memory) Create(t *tank.Tank) error {
	// Primary key..
	m.count++

	t.Id = m.count

	// Save operation
	m.mutex.Lock()
	m.records[m.count] = t
	m.mutex.Unlock()

	return nil
}

// Read tank from memory
func (m *memory) Read(id int64) (*tank.Tank, error) {
	t, ok := m.records[id]
	if !ok {
		// Not found
		return nil, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	// Search for fish
	f, err := dbFish.Search(t.Id)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// Attach to generate response
	t.Fish = f

	return t, nil
}

// Delete tank from memory
func (m *memory) Delete(id int64) error {
	t, ok := m.records[id]
	if !ok {
		// Not found
		return errors.New(strconv.Itoa(http.StatusNotFound))
	}

	// Check if tank contains fish
	f, err := dbFish.Search(t.Id)
	if err != nil {
		return errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	// There are fish, the operation is forbidden for the resource
	if len(f) != 0 {
		return errors.New(strconv.Itoa(http.StatusForbidden))
	}

	// Delete primary key from map, no-op if does not exists
	m.mutex.Lock()
	delete(m.records, id)
	m.mutex.Unlock()

	return nil
}
