package inmemory

import (
	dbTank "github.com/Rakanixu/brandcrumb/db/tank"
	// We would implement a mock to use here, and remove all depenedencies. eg)
	// _ "github.com/Rakanixu/brandcrumb/db/tank/mock"
	_ "github.com/Rakanixu/brandcrumb/db/tank/inmemory"
	"github.com/Rakanixu/brandcrumb/models/fish"
	"github.com/Rakanixu/brandcrumb/models/tank"
	"testing"
)

func TestCreate(t *testing.T) {
	m := &memory{
		records: make(map[int64]*fish.Fish),
	}

	dbTank.Create(&tank.Tank{
		Name:        "T1",
		Temperature: 20.0,
	})

	var createTestData = []struct {
		fish   *fish.Fish
		result interface{}
	}{
		// Tank exists
		{
			&fish.Fish{
				Tank:     1,
				Name:     "F1",
				FishType: "T1",
				Price:    10.99,
			},
			nil,
		},
		//Tank does not exits
		{
			&fish.Fish{
				Tank:     0,
				Name:     "F1",
				FishType: "T1",
				Price:    10.99,
			},
			"404",
		},
	}

	for _, tt := range createTestData {
		result := m.Create(tt.fish)

		if tt.result == nil {
			if tt.result != result {
				t.Errorf("Expected %v, got %v", tt.result, result)
			}
		} else {
			if tt.result != result.Error() {
				t.Errorf("Expected %v, got %v", tt.result, result.Error())
			}
		}

	}
}

func TestRead(t *testing.T) {
	m := &memory{
		records: make(map[int64]*fish.Fish),
	}

	dbTank.Create(&tank.Tank{
		Name:        "T1",
		Temperature: 20.0,
	})

	f := &fish.Fish{
		Tank:     1,
		Name:     "F1",
		FishType: "T1",
		Price:    10.99,
	}
	m.Create(f)

	var readTestData = []struct {
		fish   *fish.Fish
		result interface{}
	}{
		{
			f,
			f,
		},
	}

	for _, tt := range readTestData {
		result, _ := m.Read(tt.fish.Id)

		if tt.result != result {
			t.Errorf("Expected %v, got %v", tt.result, result)
		}

	}
}

func TestDelete(t *testing.T) {
	m := &memory{
		records: make(map[int64]*fish.Fish),
	}

	dbTank.Create(&tank.Tank{
		Name:        "T1",
		Temperature: 20.0,
	})

	f := &fish.Fish{
		Tank:     1,
		Name:     "F1",
		FishType: "T1",
		Price:    10.99,
	}
	m.Create(f)

	var deleteTestData = []struct {
		fish   *fish.Fish
		result interface{}
	}{
		{
			f,
			nil,
		},
	}

	for _, tt := range deleteTestData {
		result := m.Delete(tt.fish.Id)

		if tt.result != result {
			t.Errorf("Expected %v, got %v", tt.result, result)
		}
	}
}
