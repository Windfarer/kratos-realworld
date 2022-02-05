package server

import (
	"encoding/json"
	"fmt"
	"kratos-realworld/internal/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPStruct(t *testing.T) {
	a := &errors.HTTPError{
		Errors: make(map[string][]string),
	}
	a.Errors["body"] = []string{"can't be empty"}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Printf("%s", string(b))
}
