package common

import "errors"

var (
	ErrNoUser        = errors.New("no user found")
	ErrEmailNotValid = errors.New("email not valid")
)
