package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeathCheck(t *testing.T) {

	path, err := os.Getwd()

	assert.Nil(t, err)
	assert.NotEmpty(t, path)

	assert.Empty(t, Healthcheck(path))

}
