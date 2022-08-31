package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTime(t *testing.T) {
	t.Run("Checking for errors..", func(t *testing.T) {
		if _, err := getTime(); err != nil {
			assert.NoError(t, err)
		}
	})
}
