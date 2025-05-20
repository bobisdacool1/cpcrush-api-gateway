package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustLoadConfig(t *testing.T) {
	_, err := newConfig()
	assert.NoError(t, err)
}
