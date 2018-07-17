// Package props exposes structs and functions for manipulating of properties.
// properties are values that can be passed around to components.
package props

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gernest/vected/lib/html/template"
)

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

// Filter returns a new Props with only values that the filter function fn
// evaluates to true.
func (p Props) Filter(fn func(k, v interface{}) bool) Props {
	m := make(Props)
	for k, v := range p {
		if fn(k, v) {
			m[k] = v
		}
	}
	return m
}

// Attr returns a string for prop attributes. This only collect string keys. For
// boolean values the attribute will not contain the value part. Other types of
// kys/values are simply ignored.
//
// 	p["checked"]=true => checked
// 	p["onClick"]="handleClick" => onClick="handleClick"
func (p Props) Attr() template.HTMLAttr {
	var keys []string
	for k := range p {
		if s, ok := k.(string); ok {
			keys = append(keys, s)
		}
	}
	sort.Strings(keys)
	var o []string
	for _, k := range keys {
		v := p[k]
		switch e := v.(type) {
		case string:
			o = append(o, fmt.Sprintf(`%v="%s"`, k, e))
		case bool:
			o = append(o, fmt.Sprint(k))
		}
	}
	return template.HTMLAttr(strings.Join(o, " "))
}

// Int calls Int with p as first argument.
func (p Props) Int(key interface{}) NullInt {
	return Int(p, key)
}

// StringV calls StringV with p as first argument.
func (p Props) StringV(key interface{}, value string) NullString {
	return StringV(p, key, value)
}

// Int looks for property value with key, and tries to cast it to an int. This
// will set NullInt.IsNull to true if the key is missing or the value is not of
// type it.
func Int(p Props, key interface{}) NullInt {
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

// String tries to return the value stored by key as a string.
func (p Props) String(key interface{}) NullString {
	return String(p, key)
}

// String finds key's value in p and casts it as a string.
func String(p Props, key interface{}) NullString {
	if v, ok := p[key]; ok {
		if vi, ok := v.(string); ok {
			return NullString{Value: vi}
		}
	}
	return NullString{IsNull: true}
}

// Bool tries to find key's value in p and casts it to bool if found.
func (p Props) Bool(key interface{}) NullBool {
	return Bool(p, key)
}

// Bool finds key's value in p and casts it as a bool.
func Bool(p Props, key interface{}) NullBool {
	if v, ok := p[key]; ok {
		if vi, ok := v.(bool); ok {
			return NullBool{Value: vi}
		}
	}
	return NullBool{IsNull: true}
}
