package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	// Test data
	testData := []any{1, 2, 3}

	// Create a collection
	c := New(testData)

	assert.NotNil(t, c, "Expected collection to be created")
	assert.Equal(t, testData, c.data, "Expected collection data to match input data")
}

func TestFilter(t *testing.T) {
	// Test data
	testData := []any{1, 2, 3, 4, 5}

	// Create a collection
	c := New(testData)

	// Filter even numbers
	filtered := c.Filter(func(item any) bool {
		if num, ok := item.(int); ok {
			return num%2 == 0
		}
		return false
	})

	// Assert the filtered data
	expected := []any{2, 4}
	assert.Equal(t, expected, filtered.data, "Expected filtered data to match")
}
