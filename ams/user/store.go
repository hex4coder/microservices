package main

import "context"

type store struct {
	// database here
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(ctx context.Context) error {
	return nil
}

// Delete implements UserStore.
func (s *store) Delete(context.Context) error {
	return nil
}

// Get implements UserStore.
func (s *store) Get(context.Context) error {
	return nil
}

// Update implements UserStore.
func (s *store) Update(context.Context) error {
	return nil
}
