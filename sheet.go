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

func (s *Sheet) AddRule(rules CSSRule) {
	v := process(rules, classNamer(
		namerFunc(s.CLasses, s.idGen),
	))
	if ls, ok := v.(RuleList); ok {
		s.rules = append(s.rules, ls...)
	} else {
		s.rules = append(s.rules, v)
	}
}

func (s *Sheet) Text() string {
	return toString(s.rules)
}

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
		o[v] = true
	}
	return o
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
				switch ne := v.(type) {
				case StyleRule:
					if isClass(ne.Selector) {
						ne.Selector = namer(ne.Selector)
					}
					o = append(o, ne)
				case RuleList:
					for _, value := range ne {
						if st, ok := value.(StyleRule); ok {
							if isClass(st.Selector) {
								st.Selector = namer(st.Selector)
							}
						}
						o = append(o, value)
					}
				default:
					o = append(o, ne)
				}
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
		s.list = append(s.list, toString(v, Options{NoPretty: true}))
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
