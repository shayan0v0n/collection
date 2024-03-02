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
