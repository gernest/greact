package gs

type Sheet struct {
	id       int64
	CLasses  ClassMap
	rules    RuleList
	list     []string
	idGen    func() string
	attached bool
	registry Registry
}

// AddRule processes the rules are returns mapping of original classnames to the
// generated classnames.
//
// The resulting class map is merged with the existing classmap, so added
// classes can always be acced by s.Classes[className] field. The processed
// rules are stored in the stylesheet, note that this doesn't attach the sheet
// to the dom. You need to explicitly call Attach method to attach the styles to
// the dom.
func (s *Sheet) AddRule(rules CSSRule) ClassMap {
	m := make(ClassMap)
	v := Process(rules, classNamer(
		namerFunc(m, s.idGen),
	))
	if ls, ok := v.(RuleList); ok {
		s.rules = append(s.rules, ls...)
	} else {
		s.rules = append(s.rules, v)
	}
	for k, v := range m {
		s.CLasses[k] = v
	}
	return m
}

// func (s *Sheet) Text() string {
// 	return s.rules.String()
// }

func NewSheet(idGen func() string) *Sheet {
	return &Sheet{CLasses: make(ClassMap), idGen: idGen}
}

func namerFunc(c ClassMap, idGen func() string) func(string) string {
	return func(s string) string {
		if exist, ok := c[s]; ok {
			return exist
		}
		gen := s + "-" + idGen()
		c[s] = gen
		return gen
	}
}

type ClassMap map[string]string

// Classes to be compatible with vecty.ClassMap
func (c ClassMap) Classes() map[string]bool {
	o := make(map[string]bool)
	for _, v := range c {
		o[toClassname(v)] = true
	}
	return o
}
func toClassname(c string) string {
	if c == "" {
		return c
	}
	if c[0] == '.' {
		return c[1:]
	}
	return c
}
func isClass(s string) bool {
	return s != "" && s[0] == '.'
}

func classNamer(namer func(string) string) Transformer {
	return func(rules CSSRule) CSSRule {
		switch e := rules.(type) {
		case StyleRule:
			if isClass(e.Selector) {
				e.Selector = namer(e.Selector)
			}
			return e
		case RuleList:
			var o RuleList
			for _, v := range e {
				o = append(o, classNamer(namer)(v))
			}
			return o
		default:
			return e
		}
	}
}

func (s *Sheet) Attach() {
	if !s.attached {
		s.registry.Attach(s)
		s.attached = true
	}
}

func (s *Sheet) Detach() {
	if s.attached {
		s.registry.Detach(s)
		s.attached = false
	}
}

func (s *Sheet) ListRules() []string {
	if s.list != nil {
		return s.list
	}
	for _, v := range s.rules {
		s.list = append(s.list, v.String())
	}
	return s.list
}

type Registry interface {
	NewSheet() *Sheet
	Attach(*Sheet)
	Detach(*Sheet)

	//again no user implementation
	isRegistry()
}
