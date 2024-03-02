package collection

import "reflect"

// Collection represents a collection of data.
type Collection struct {
	data []any
}

// New creates a new collection with the given data.
func New(data []any) *Collection {
	return &Collection{
		data: data,
	}
}

// Filter applies a filter function to the collection and returns a new filtered collection.
func (c *Collection) Filter(fn func(any) bool) *Collection {
	var filtered []any
	for _, item := range c.data {
		if fn(item) {
			filtered = append(filtered, item)
		}
	}
	return New(filtered)
}

// Map applies a transformation function to each item in the collection and returns a new collection with the transformed items.
func (c *Collection) Map(fn func(any) any) *Collection {
	var mapped []any
	for _, item := range c.data {
		mapped = append(mapped, fn(item))
	}
	return New(mapped)
}

// All returns all elements of a collection.
func (c *Collection) All() []any {
	return c.data
}

func (c *Collection) Reject(fn func(any) bool) *Collection {
	var filtered []any
	for _, item := range c.data {
		if !fn(item) {
			filtered = append(filtered, item)
		}
	}
	return New(filtered)
}

// Avg get the average value of a given key.
func (c *Collection) Avg() float64 {
	var avg float64
	if len(c.data) == 0 {
		return avg
	}
	for _, item := range c.data {
		v := reflect.ValueOf(item)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			avg += float64(v.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			avg += float64(v.Uint())
		case reflect.Float32, reflect.Float64:
			avg += v.Float()
		default:
			panic("Unsupported operand.")
		}
	}
	return avg / float64(len(c.data))
}
