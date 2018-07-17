package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"text/template"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"github.com/urfave/cli"
)

const (
	runtimeSrc = "lib/runtimejs/js"
	runtimeOut = "lib/runtimejs"
)

type jsSrc struct {
	Name     string
	OutPath  string
	Plain    string
	Minified string
}

// Runtime creates lib/include.gen.go file that contains both minified
// an un minified source for include.js library found in lib/include/include.js.
func Runtime(ctx *cli.Context) error {
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	info, err := ioutil.ReadDir(runtimeSrc)
	if err != nil {
		return err
	}
	var files []string
	for _, v := range info {
		if v.IsDir() {
			continue
		}
		if filepath.Ext(v.Name()) == ".js" {
			files = append(files, filepath.Join(runtimeSrc, v.Name()))
		}
	}
	sort.Strings(files)
	tplTxt := `package runtimejs 
	// {{.name}}Plain is a plain {{.file}} content as string
	var {{.name}}Plain ={{.plain}}

	// {{.name}}Min is a minified {{.file}} content as string
	var {{.name}}Min={{.minified}}
	`
	tpl, err := template.New("runtimejs").Parse(tplTxt)
	if err != nil {
		return err
	}
	var names []string
	for _, v := range files {
		name := strings.TrimSuffix(filepath.Base(v), filepath.Ext(v))
		names = append(names, name)
		f, err := ioutil.ReadFile(v)
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err := m.Minify("text/javascript", &buf, bytes.NewReader(f)); err != nil {
			return err
		}

		var out bytes.Buffer
		err = tpl.Execute(&out, map[string]string{
			"file":     v,
			"name":     name,
			"plain":    fmt.Sprintf("`%s`", string(f)),
			"minified": fmt.Sprintf("`%s`", buf.String()),
		})
		b, err := format.Source(out.Bytes())
		if err != nil {
			return err
		}
		o := filepath.Join(runtimeOut, fmt.Sprintf("%s.gen.go", name))
		err = ioutil.WriteFile(o, b, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

// RuntimeCMD returns a command for generating include package.
func RuntimeCMD() cli.Command {
	return cli.Command{
		Name:   "runtimejs",
		Action: Runtime,
	}
}
