package validator

import (
	"encoding/json"
	"fmt"
)

type Errors map[string][]string

func (e Errors) AddErrors(name string, errors ...string) {
	e[name] = append(e[name], errors...)
}

func (e Errors) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", data)
}
