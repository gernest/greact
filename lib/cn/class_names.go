// Package cn provide utilities for ma=handling css class names.
package cn

// Name represent a css class. If Skip is set to true this class will be
// ignored.
type Name struct {
	// Is the class name.
	C string

	// stands for SKip, if true then the name will not be included when joining.
	S bool
}

// N retruns Name struct with klass value as C and skip value as S.
func N(klass string, skip bool) Name {
	return Name{C: klass, S: skip}
}

// Join joins the names to form css classes.
func Join(names ...interface{}) string {
	c := &class{}
	for _, name := range names {
		switch t := name.(type) {
		case string:
			if t == "" {
				continue
			}
			c.add(t)
		case Name:
			if t.S || t.C == "" {
				continue
			}
			c.add(t.C)
		}
	}
	return c.String()
}

type class struct {
	n []string
}

func (c *class) add(cl string) {
	c.n = append(c.n, cl)
}

func (c *class) String() string {
	buf := ""
	for k, v := range c.n {
		if k == 0 {
			buf += v
		} else {
			buf += " " + v
		}
	}
	return buf
}
