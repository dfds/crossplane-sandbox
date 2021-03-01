package misc

import (
	v1 "k8s.io/api/core/v1"
	v1Networking "k8s.io/api/networking/v1"
)

type Store struct {
	Services map[string]v1.Service
	Ingresses map[string]v1Networking.Ingress
}

func NewStore() *Store{
	return &Store {
		Services: make(map[string]v1.Service),
		Ingresses: make(map[string]v1Networking.Ingress),
	}
}

func (s *Store) PutService(key string, val v1.Service) {
	s.Services[key] = val
}

func (s *Store) GetService(key string) v1.Service {
	return s.Services[key]
}

func (s *Store) GetServices() map[string]v1.Service {
	return s.Services
}


func (s *Store) RemoveService(key string, val v1.Service) {
	delete(s.Services, key)
}

