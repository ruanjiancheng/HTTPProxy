package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Print(t *testing.T) {
	config, err := ReadConfig("config.yaml")
	assert.Equal(t, err, nil)
	config.Print()
}
