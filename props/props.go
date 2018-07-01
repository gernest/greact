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

// Int calls Int with p as first argument.
func (p Props) Int(key interface{}) NullInt {
	return Int(key, p)
}

// StringV calls StringV with p as first argument.
func (p Props) StringV(key interface{}, value string) NullString {
	return StringV(p, key, value)
}

// Int looks for property value with key, and tries to cast it to an int. This
// will set NullInt.IsNull to true if the key is missing or the value is not of
// type it.
func Int(key interface{}, p Props) NullInt {
	if v, ok := p[key]; ok {
		if vi, ok := v.(int); ok {
			return NullInt{Value: vi}
		}
	}
	return NullInt{IsNull: true}
}

// StringV looks for key's value in ctx. This will return value if the key is
// missing, this function will try to cast the value to string.When the value is
// not of type string the NullString.IsNull field will be set to true.
func StringV(ctx Props, key interface{}, value string) NullString {
	if v, ok := ctx[key]; ok {
		if vk, ok := v.(string); ok {
			return NullString{Value: vk}
		}
		return NullString{IsNull: true}
	}
	return NullString{Value: value}
}
