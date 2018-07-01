// Package props exposes structs and functions for manipulating of properties.
// properties are values that can be passed around to components.
package props

// Props is a map of properties. These are used to pass values to components.
type Props map[interface{}]interface{}

// Merge returns new Props with values from b added to a.
func Merge(a, b Props) Props {
	m := make(Props)
	for k, v := range a {
		m[k] = v
	}
	for k, v := range b {
		m[k] = v
	}
	return m
}
