package collection

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
