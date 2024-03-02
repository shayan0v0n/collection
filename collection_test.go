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

func TestMap(t *testing.T) {
	// Test data
	testData := []any{1, 2, 3, 4, 5}

	// Create a collection
	c := New(testData)

	// Map function to double each number
	mapped := c.Map(func(item any) any {
		if num, ok := item.(int); ok {
			return num * 2
		}
		return item
	})

	// Assert the mapped data
	expected := []any{2, 4, 6, 8, 10}
	assert.Equal(t, expected, mapped.data, "Expected mapped data to match")
}

func TestAll(t *testing.T) {
	// Create a collection with some test data
	testData := []interface{}{1, 2, 3, 4, 5}
	c := New(testData)

	// Call All method
	allData := c.All()
	filteredData := c.Filter(func(item any) bool {
		return item.(int)%2 == 0
	}).All()

	assert.Equal(t, allData, testData)
	assert.Equal(t, []any{2, 4}, filteredData)
}
