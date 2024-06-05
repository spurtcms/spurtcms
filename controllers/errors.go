package controllers

import "errors"

var (
	ErrInternalServerError = "internalserverr"
	ErrLoadLocation        = errors.New("timeZone not found")
)
