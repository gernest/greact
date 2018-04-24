package gs

type Sheet struct {
	ID      string
	CLasses ClassMap
	Text    string
}

func NewSheet(rules CSSRule, idGen func() string) *Sheet {
	s := &Sheet{CLasses: make(ClassMap)}
	v := ToString(rules, classNamer(
		namerFunc(s.CLasses, idGen),
	))
	s.Text = v
	return s
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
