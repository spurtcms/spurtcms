package controllers

import "errors"

var (
	ErrInternalServerError       = "internalserverr"
	ErrLoadLocation              = errors.New("timeZone not found")
	ErrGettingTemplateList       = errors.New("unable to fetch the template list")
	ErrGettingTemplateModuleList = errors.New("unable to fetch the template module list")
)
