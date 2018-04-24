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
		name := s[1:]
		gen := s + "-" + idGen()
		c[name] = gen
		return gen
	}
}

type ClassMap map[string]string

func isClass(s string) bool {
	return s != "" && s[0] == '.'
}

func classNamer(namer func(string) string) Transformer {
	return func(rules CSSRule) CSSRule {
		switch e := rules.(type) {
		case style:
			if isClass(e.selector) {
				e.selector = namer(e.selector)
			}
			return e
		case RuleList:
			var o RuleList
			for _, v := range e {
				switch ne := v.(type) {
				case style:
					if isClass(ne.selector) {
						ne.selector = namer(ne.selector)
					}
					o = append(o, ne)
				case RuleList:
					for _, value := range ne {
						if st, ok := value.(style); ok {
							if isClass(st.selector) {
								st.selector = namer(st.selector)
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
