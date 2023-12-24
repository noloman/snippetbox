package models

import (
	"errors"
)

var (
	// ErrNoRecord adds a new ErrNoRecord error. We'll use this later if a record can't be found
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials adds a new ErrInvalidCredentials error. We'll use this later if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail adds a new ErrDuplicateEmail error. We'll use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
