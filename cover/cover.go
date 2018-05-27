package cover

import (
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/gernest/mad/config"
	"github.com/mafredri/cdp/protocol/profiler"
	sourcemap "gopkg.in/sourcemap.v1"
)

// SourceFile stores coverage profile detail of a file.
type SourceFile struct {
	Name   string      `json:"name"`
	Source string      `json:"source"`
	Lines  map[int]int `json:"lines"`
}

// Process calculates code coverage  from the chrome coverage profiles.
func Process(cfg *config.Config, profiles []profiler.Profile) error {
	sourceMaps := make(map[string]*sourcemap.Consumer)
	sourceFiles := make(map[string]*SourceFile)
	for _, v := range profiles {
		for _, node := range v.Nodes {
			if strings.Contains(node.CallFrame.FunctionName, cfg.Info.ImportPath) {
				c, ok := sourceMaps[node.CallFrame.URL]
				if !ok {
					u, err := url.Parse(node.CallFrame.URL)
					if err != nil {
						return err
					}
					src := u.Query().Get("src")
					src, err = url.QueryUnescape(src)
					if err != nil {
						return err
					}
					path := filepath.Join(cfg.OutputPath, src)
					dir := filepath.Dir(path)
					base := filepath.Base(path)
					jsmap := filepath.Join(dir, base+".map")
					data, err := ioutil.ReadFile(jsmap)
					if err != nil {
						return err
					}
					smap, err := sourcemap.Parse(base, data)
					if err != nil {
						return err
					}
					sourceMaps[node.CallFrame.URL] = smap
					c = smap
				}
				line := node.CallFrame.LineNumber
				col := node.CallFrame.ColumnNumber
				name, _, line, _, ok := c.Source(line, col)
				if strings.Contains(name, cfg.Info.ImportPath) {
					if ok {
						s, ok := sourceFiles[name]
						if !ok {
							fn := fileName(cfg, name)
							file := filepath.Join(cfg.Info.Dir, fn)
							data, err := ioutil.ReadFile(file)
							if err != nil {
								return err
							}
							s = &SourceFile{
								Name:   fn,
								Source: string(data),
								Lines:  make(map[int]int),
							}
							sourceFiles[name] = s
						}
						count := 0
						if node.HitCount != nil {
							count = *node.HitCount
						}
						if n, ok := s.Lines[line]; ok {
							s.Lines[line] = n + count
						} else {
							s.Lines[line] = count
						}
					}
				}
			}
		}
	}
	return nil
}

func fileName(cfg *config.Config, path string) string {
	path = strings.TrimPrefix(path, string(filepath.Separator))
	path = strings.TrimPrefix(path, cfg.Info.ImportPath)
	path = strings.TrimPrefix(path, string(filepath.Separator))
	return path
}
