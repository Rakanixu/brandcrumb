package inmemory

import (
	// We would implement a mock to use here, and remove all depenedencies. eg)
	// _ "github.com/Rakanixu/brandcrumb/db/fish/mock"
	_ "github.com/Rakanixu/brandcrumb/db/fish/inmemory"
	"github.com/Rakanixu/brandcrumb/models/tank"
	"testing"
)

func TestCreate(t *testing.T) {
	m := &memory{
		records: make(map[int64]*tank.Tank),
	}

	var createTestData = []struct {
		tank   *tank.Tank
		result interface{}
	}{
		{
			&tank.Tank{
				Name:        "T1",
				Temperature: 20.5,
			},
			nil,
		},
	}

	for _, tt := range createTestData {
		result := m.Create(tt.tank)

		if tt.result != result {
			t.Errorf("Expected %v, got %v", tt.result, result)
		}
	}
}

func TestRead(t *testing.T) {
	m := &memory{
		records: make(map[int64]*tank.Tank),
	}

	t1 := &tank.Tank{
		Name:        "T1",
		Temperature: 20.5,
	}
	m.Create(t1)

	var readTestData = []struct {
		tank *tank.Tank
	}{
		{
			t1,
		},
	}

	for _, tt := range readTestData {
		result, _ := m.Read(tt.tank.Id)

		if tt.tank != result {
			t.Errorf("Expected %v, got %v", tt.tank, result)
		}

	}
}

func TestDelete(t *testing.T) {
	m := &memory{
		records: make(map[int64]*tank.Tank),
	}

	t1 := &tank.Tank{
		Name:        "T1",
		Temperature: 20.5,
	}
	m.Create(t1)

	var deleteTestData = []struct {
		tank   *tank.Tank
		result interface{}
	}{
		{
			t1,
			nil,
		},
	}

	for _, tt := range deleteTestData {
		result := m.Delete(tt.tank.Id)

		if tt.result != result {
			t.Errorf("Expected %v, got %v", tt.result, result)
		}
	}
}
