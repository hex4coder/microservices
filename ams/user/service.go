package main

import "context"

type service struct {
	store UserStore
}

func NewService(store UserStore) *service {
	return &service{store: store}
}

func (s *service) CreateUser(ctx context.Context) error {
	return nil
}
