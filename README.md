# Go Collection Package

The Go Collection Package provides a collection data structure similar to Laravel Collection in PHP. It allows you to
work with arrays of data in a fluent and expressive manner, providing methods for filtering, transforming, and
manipulating data.

## Installation

To use the collection package in your Go project, you can install it using:

```bash
go get github.com/kamandlou/collection
```

## List of methods
- Filter
- Map
- All
- Reject
- Avg

## Usage
```go
data := []any{1, 2, 3, 4, 5}

// Create a collection
c := collection.New(data)

// Filter even numbers
filteredData := c.Filter(func(item any) bool {
    if num, ok := item.(int); ok {
        return num%2 == 0
    }
    return false
}).All()

fmt.Println(filteredData) // []any{2, 4}
```

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on GitHub.

We are trying to implement these methods on this project. [list of methods](https://laravel.com/docs/10.x/collections#available-methods)


## License
This project is licensed under the Apache License - see the [LICENSE](./LICENSE) file for details.
