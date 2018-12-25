package gen

import (
	"bytes"
	"go/format"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
	"golang.org/x/net/html/atom"
)

const mozAttributeReference = "https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes"

func downloadAttrs(dest string) error {
	res, err := http.Get(mozAttributeReference)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, b, 0600)
}

type attribute struct {
	Name     string
	Elements []string
}

var elRegexp = regexp.MustCompile(`[\<\>\s]`)

func extractAttr(doc *goquery.Document) ([]attribute, error) {
	var rst []attribute
	table := doc.Find(".standard-table")
	table.Find("tbody tr").Each(func(i int, sel *goquery.Selection) {
		children := sel.Find("td")
		a := children.Eq(0).Text()
		e := children.Eq(1).Text()
		e = elRegexp.ReplaceAllString(e, "")
		parts := strings.Split(e, ",")
		rst = append(rst, attribute{
			Name:     a,
			Elements: parts,
		})
	})
	return rst, nil
}

const astr = `package attribute
// Attribute represents htm attributes
type Attribute struct{
	Name string
	Elements []string
}
// Map maps html attribute name to the Attribute object.
var Map =map[string]Attribute{
	{{- range $k,$v:=. }}
	"{{.Name}}": Attribute{
		Name:"{{.Name}}",
		{{- with .Elements}}
		Elements: []string{
			{{- range . -}}
			"{{.}}",
			{{- end}}
		},
		{{- end}}
	},
	{{- end -}}
}
`

func generateAttributes(w io.Writer, a []attribute) error {
	tpl, err := template.New("a").Funcs(template.FuncMap{
		"toAtom": toAtom,
		"len": func(a []attribute) int {
			return len(a) - 1
		},
	}).Parse(astr)
	if err != nil {
		return err
	}
	return tpl.Execute(w, a)
}

func toAtom(e string) string {
	return "atom." + strings.Title(atom.Lookup([]byte(e)).String())
}

func AttrCMD() cli.Command {
	return cli.Command{
		Name:  "attr",
		Usage: "generates html attributes package",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "out",
				Value: "attribute/attributes.go",
			},
		},
		Action: func(ctx *cli.Context) error {
			out := ctx.String("out")
			var buf bytes.Buffer
			q, err := goquery.NewDocument(mozAttributeReference)
			if err != nil {
				return err
			}
			a, err := extractAttr(q)
			if err != nil {
				return err
			}
			err = generateAttributes(&buf, a)
			if err != nil {
				return err
			}
			b, err := format.Source(buf.Bytes())
			if err != nil {
				return err
			}
			return ioutil.WriteFile(out, b, 0600)
		},
	}
}
