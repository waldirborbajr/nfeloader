package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSN(t *testing.T) {

	assert.NotEmpty(t, Dsn())

}
