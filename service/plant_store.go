package service

import (
	"errors"
	"fmt"
	"golang_grpc_unary_stream/pb"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("[ERROR] Plant record already exist")

type PlantStore interface {
	// Store the plant data into the plant store
	Save(plant *pb.Plant) error
	// Fetch the plant data from the plant store using plant id
	Find(id string) (*pb.Plant, error)
}

type InMemoryPlantStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Plant
}

// NewInMemoryPlantStore used to create new in memory plant store
func NewInMemoryPlantStore() *InMemoryPlantStore {
	return &InMemoryPlantStore{
		data: make(map[string]*pb.Plant),
	}
}

// Save used to insert the plant data into the plant store
func (store *InMemoryPlantStore) Save(plant *pb.Plant) error {

	// Lock locks store.
	// If the lock is already in use, the calling goroutine blocks until the mutex is available.
	store.mutex.Lock()
	// Unlock unlocks store.
	// It is a run-time error if m is not locked on entry to Unlock.
	defer store.mutex.Unlock()

	if store.data[plant.Id] != nil {
		return ErrAlreadyExists
	}

	other := &pb.Plant{}
	err := copier.Copy(other, plant)
	if err != nil {
		return fmt.Errorf("[ERROR] While, Copying plant data: %v", err)
	}
	store.data[other.Id] = other
	return nil

}

// Find used to fetch the plant data from plant store using plant id
func (store *InMemoryPlantStore) Find(id string) (*pb.Plant, error) {

	// RLock locks store for reading.
	store.mutex.RLock()
	// RLock unlocks store for reading.
	defer store.mutex.RUnlock()

	plant := store.data[id]
	if plant == nil {
		return nil, nil
	}

	other := &pb.Plant{}
	err := copier.Copy(other, plant)
	if err != nil {
		return nil, fmt.Errorf("[ERRORO] While, Copying plant data: %v", err)
	}

	return other, nil
}
