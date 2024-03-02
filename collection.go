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
	return &Collection{data: filtered}
}
