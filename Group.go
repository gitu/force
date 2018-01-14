package main

import (
	"gopkg.in/validator.v2"
)

type Group struct {
	Id   string `json:"id" validate:"nonzero"`
	Name string `json:"name" validate:"nonzero"`
}

func (s Group) Validate() error {

	return validator.Validate(s)
}
