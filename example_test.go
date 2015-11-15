package gstring

import (
	"math"
)

func Example_basic() {
	Printm("{test}", map[string]interface{}{"test": "hello world"})

	// Output: hello world
}

func Example_formatting() {
	Printm("{pi:%2.5f}", map[string]interface{}{"pi": math.Pi})

	// Output: 3.14159
}
