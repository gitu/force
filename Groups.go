package main

import (
	"fmt"
)

type Groups []Group

func (s Groups) Validate() error {

	// make sure it is unique
	m := map[interface{}]struct{}{}
	for _, v := range s {
		m[v] = struct{}{}
	}
	if len(m) != len(s) {
		return fmt.Errorf("collection is not unique")
	}

	return nil
}
