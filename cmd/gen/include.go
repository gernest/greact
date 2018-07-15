package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"

	"text/template"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"github.com/urfave/cli"
)

const (
	includeSrc    = "lib/runtimejs/include.js"
	includeOutput = "lib/runtimejs/include.gen.go"
)

// GenerateInclude creates lib/include.gen.go file that contains both minified
// an un minified source for include.js library found in lib/include/include.js.
func GenerateInclude(ctx *cli.Context) error {
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	f, err := ioutil.ReadFile(includeSrc)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := m.Minify("text/javascript", &buf, bytes.NewReader(f)); err != nil {
		return err
	}
	tplTxt := `package runtime 
	// Plain is a plain include.js content as string
	var IncludeMin ={{.plain}}

	// IncludeMin is a minified include.js content as string
	var IncludeMin={{.minified}}
	`
	tpl, err := template.New("include").Parse(tplTxt)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	err = tpl.Execute(&out, map[string]string{
		"plain":    fmt.Sprintf("`%s`", string(f)),
		"minified": fmt.Sprintf("`%s`", buf.String()),
	})
	b, err := format.Source(out.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(includeOutput, b, 0600)
}

// Include returns a command for generating include package.
func Include() cli.Command {
	return cli.Command{
		Name:   "include",
		Action: GenerateInclude,
	}
}
