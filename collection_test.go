package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	testData := []any{1, 2, 3}
	c := New(testData)

	assert.NotNil(t, c, "Expected collection to be created")
	assert.Equal(t, testData, c.data, "Expected collection data to match input data")
}
