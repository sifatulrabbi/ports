package services

import (
	"fmt"

	"gorm.io/gorm"
)

var _ crudService[interface{}, interface{}, interface{}] = &ServiceWithDB[interface{}, interface{}, interface{}]{}

type crudService[Data any, Payload any, Filter any] interface {
	Log(message string)
	CreateOne(p Payload) (*Data, error)
	UpdateOne(p Payload) (*Data, error)
	DeleteOne(f Filter) (*Data, error)
	GetOne(f Filter) (*Data, error)
	GetMany(f Filter) (*Data, error)
}

type ServiceWithDB[Data any, Payload any, Filter any] struct {
	db   *gorm.DB
	name string
}

func (s *ServiceWithDB[Data, Payload, Filter]) Log(message string) {
	fmt.Printf("[%s] %s\n", s.name, message)
}

func (s *ServiceWithDB[Data, Payload, Filter]) CreateOne(p Payload) (*Data, error) {
	return nil, fmt.Errorf("service isn't instantiated properly")
}

func (s *ServiceWithDB[Data, Payload, Filter]) UpdateOne(p Payload) (*Data, error) {
	return nil, fmt.Errorf("service isn't instantiated properly")
}

func (s *ServiceWithDB[Data, Payload, Filter]) DeleteOne(f Filter) (*Data, error) {
	return nil, fmt.Errorf("service isn't instantiated properly")
}

func (s *ServiceWithDB[Data, Payload, Filter]) GetOne(f Filter) (*Data, error) {
	return nil, fmt.Errorf("service isn't instantiated properly")
}

func (s *ServiceWithDB[Data, Payload, Filter]) GetMany(f Filter) (*Data, error) {
	return nil, fmt.Errorf("service isn't instantiated properly")
}

func newServiceWithDB[Data any, Payload any, Filter any](db *gorm.DB, serviceName string) *ServiceWithDB[Data, Payload, Filter] {
	return &ServiceWithDB[Data, Payload, Filter]{
		db:   db,
		name: serviceName,
	}
}
