package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")
var ErrInvalidCredentials = errors.New("models: invalid credentials")
var ErrDuplicateEmail = errors.New("models: duplicate email")
var ErrInvalidInput = errors.New("models: invalid input")
var ErrUnauthorized = errors.New("models: unauthorized access")
var ErrInternalServer = errors.New("models: internal server error")
