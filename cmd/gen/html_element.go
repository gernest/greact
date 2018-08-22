package gen

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"regexp"
	"sort"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
)

const mozElementReference = "https://developer.mozilla.org/en-US/docs/Web/HTML/Element"

var re = regexp.MustCompile(`[\<\>]`)

func extractElements(doc *goquery.Document) ([]byte, error) {
	sel := doc.Find("td:first-child").Map(func(i int, s *goquery.Selection) string {
		return re.ReplaceAllString(s.Text(), "")
	})
	add := []string{
		"base",
		"section",
		"h1",
		"h2",
		"h3",
		"h4",
		"h5",
		"h6",
		"iframe",
	}
	sel = append(sel, add...)
	var ctx []string
	m := make(map[string]bool)
	for _, v := range sel {
		if !m[v] {
			ctx = append(ctx, v)
			m[v] = true
		}
	}
	sort.Strings(ctx)

	tpl, err := template.New("el").Parse(`package elements

	var elems =map[string]bool{
		{{- range .}}
		"{{.}}":true,
		{{- end}}
	}
	// Valid returns true if the name is a valid html element
	func Valid(name string)bool  {
		return elems[name]
	}
	`)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, ctx)
	if err != nil {
		return nil, err
	}
	return format.Source(buf.Bytes())
}

func ElementsCMD() cli.Command {
	return cli.Command{
		Name:  "elems",
		Usage: "generates html elements package",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "out",
				Value: "elements/elements.go",
			},
		},
		Action: func(ctx *cli.Context) error {
			out := ctx.String("out")
			q, err := goquery.NewDocument(mozElementReference)
			if err != nil {
				return err
			}
			a, err := extractElements(q)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(out, a, 0600)
		},
	}
}
