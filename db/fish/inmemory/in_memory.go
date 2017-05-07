package inmemory

import (
	"errors"
	dbFish "github.com/Rakanixu/brandcrumb/db/fish"
	dbTank "github.com/Rakanixu/brandcrumb/db/tank"
	"github.com/Rakanixu/brandcrumb/models/fish"
	"net/http"
	"strconv"
	"sync"
)

// memory holds all fish in the system
// it is an approach to a relational DB, so this represents the fish table withing its operations
type memory struct {
	mutex   sync.Mutex
	records map[int64]*fish.Fish
	count   int64
}

// init auto register implementation when imported
func init() {
	dbFish.Register(&memory{
		records: make(map[int64]*fish.Fish),
	})
}

// Create fish in memory
func (m *memory) Create(f *fish.Fish) error {
	// Check foreign key, drop fish if tank does not exists
	t, err := dbTank.Read(f.Tank)
	if err != nil {
		return err
	}

	// Tank already exists
	if t != nil {
		// Assign primary key
		m.count++

		f.Id = m.count

		// Save operation
		m.mutex.Lock()
		m.records[m.count] = f
		m.mutex.Unlock()
	}

	return nil
}

// Read fish from memory
func (m *memory) Read(id int64) (*fish.Fish, error) {
	t, ok := m.records[id]
	if !ok {
		// Not found
		return nil, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	return t, nil
}

// Delete fish from memory
func (m *memory) Delete(id int64) error {
	_, ok := m.records[id]
	if !ok {
		// Not found
		return errors.New(strconv.Itoa(http.StatusNotFound))
	}

	// Delete primary key from map
	m.mutex.Lock()
	delete(m.records, id)
	m.mutex.Unlock()

	return nil
}

// Search fish from memory, returns all fish beloging to the tank id
func (m *memory) Search(id int64) ([]*fish.Fish, error) {
	var result []*fish.Fish

	for _, v := range m.records {
		if v.Tank == id {
			result = append(result, v)
		}
	}

	return result, nil
}
