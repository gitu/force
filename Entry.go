package main

import (
	"github.com/gitu/force/goraml"
	"gopkg.in/validator.v2"
)

type Entry struct {
	End     goraml.DateTime `json:"end" validate:"nonzero"`
	GroupId string          `json:"groupId" validate:"nonzero"`
	Id      string          `json:"id" validate:"nonzero"`
	Name    string          `json:"name" validate:"nonzero"`
	Start   goraml.DateTime `json:"start" validate:"nonzero"`
	Type    EnumEntryType   `json:"type" validate:"nonzero"`
}

func (s Entry) Validate() error {

	return validator.Validate(s)
}
