package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/jirawan-chuapradit/grpc-golang-pcbook/pb"
)

// ErrAlreadyExists is returned when record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save saves the laptop to the store
	Save (laptop *pb.Laptop) error
	// Find finds a laptop by ID
	Find(id string) (*pb.Laptop, error)
}

// InMemoryLaptopStore store laptop in memory
type InMemoryLaptopStore struct {
	mutext sync.RWMutex
	data map[string]*pb.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}	
}

func (store *InMemoryLaptopStore) Save (laptop *pb.Laptop) error{
	store.mutext.Lock()
	defer store.mutext.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other
	return nil
}

	
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error){
	store.mutext.RLock()
	defer store.mutext.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}