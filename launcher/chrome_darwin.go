package launcher

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var priorities = []*priority{
	{
		regex:  regexp.MustCompile(fmt.Sprintf(`^%s/Applications/.*Chrome.app`, os.Getenv("HOME"))),
		weight: 50,
	},
	{
		regex:  regexp.MustCompile(fmt.Sprintf(`^%s/Applications/.*Chrome Canary.app`, os.Getenv("HOME"))),
		weight: 51,
	},
	{
		regex:  regexp.MustCompile(`^\/Applications\/.*Chrome.app`),
		weight: 100,
	},
	{
		regex:  regexp.MustCompile(`^\/Applications\/.*Chrome Canary.app`),
		weight: 101,
	},
	{
		regex:  regexp.MustCompile(`^\/Volumes\/.*Chrome.app`),
		weight: -2,
	},
	{
		regex:  regexp.MustCompile(`^\/Volumes\/.*Chrome Canary.app/`),
		weight: -1,
	},
}

func resolveChromePath() ([]string, error) {
	suffixes := []string{
		"/Contents/MacOS/Google Chrome Canary",
		"/Contents/MacOS/Google Chrome",
	}
	var install []string
	s := resolve()
	if s != "" {
		install = append(install, s)
	}
	ls := "/System/Library/Frameworks/CoreServices.framework" +
		"/Versions/A/Frameworks/LaunchServices.framework" +
		"/Versions/A/Support/lsregister"
	x := fmt.Sprintf(`%s  -dump| grep -i "google chrome\\( canary\\)\\?.app$" | awk '{$1=""; print $0}'`, ls)
	fmt.Println(x)
	o, err := exec.Command("bash", "-c", x).Output()
	if err != nil {
		return nil, fmt.Errorf("%s %v", string(o), err)
	}
	str := strings.TrimSpace(string(o))
	if str != "" {
		scan := bufio.NewScanner(strings.NewReader(str))
		scan.Split(bufio.ScanLines)
		for scan.Scan() {
			txt := scan.Text()
			txt = strings.TrimSpace(txt)
			for _, v := range suffixes {
				execPath := filepath.Join(txt, v)
				_, err := os.Stat(execPath)
				if err == nil {
					install = append(install, execPath)
				}
			}
		}
	}
	if os.Getenv(chromePath) != "" {
		priorities = append([]*priority{
			{regex: regexp.MustCompile(os.Getenv(chromePath)),
				weight: 150},
		}, priorities...)
	}
	return sortStuff(install, priorities), nil
}

func resolve() string {
	p := os.Getenv(chromePath)
	_, err := os.Stat(p)
	if err != nil {
		return ""
	}
	return p
}
