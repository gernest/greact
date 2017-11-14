package goss

type CSS map[string]interface{}

//go:generate go run bin/cssprops/main.go
func ParseCSS(css CSS) (*Style, error) {
	return nil, nil
}
