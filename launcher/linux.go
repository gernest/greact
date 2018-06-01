package launcher

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
)

var errChromePathNotSet = errors.New("the environment variable CHROME_PATH must be set to executable of a build of Chromium version 54.0 or later")

func resolveChromePathLinux() ([]string, error) {
	var install []string
	c := resolve()
	if c != "" {
		install = append(install, c)
	}
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	desktonInstall := []string{
		filepath.Join(u.HomeDir, ".local/share/applications/"),
		"/usr/share/applications",
	}
	for _, v := range desktonInstall {
		i, _ := findChromeExecutables(v)
		if i != nil {
			install = append(install, i...)
		}
	}
	executables := []string{
		"google-chrome-stable", "google-chrome",
		"chromium-browser", "chromium",
	}
	for _, v := range executables {
		o, err := exec.Command("which", v).Output()
		if err != nil {
			continue
		}
		str := string(o)
		_, err = os.Stat(str)
		if err != nil {
			continue
		}
		install = append(install, str)
	}
	if len(install) == 0 {
		return nil, errChromePathNotSet
	}
	priorities := []priority{
		{
			regex:  regexp.MustCompile(`/chrome-wrapper$/`),
			weight: 51,
		},
		{
			regex:  regexp.MustCompile(`/google-chrome-stable$/`),
			weight: 50,
		},
		{
			regex:  regexp.MustCompile(`/google-chrome$/`),
			weight: 49,
		},
		{
			regex:  regexp.MustCompile(`/chromium-browser$/`),
			weight: 48,
		},
		{
			regex:  regexp.MustCompile(`/chromium$/`),
			weight: 42,
		},
	}
	if os.Getenv(chromePath) != "" {
		priorities = append([]priority{
			{regex: regexp.MustCompile(os.Getenv(chromePath)),
				weight: 150},
		}, priorities...)
	}
	return sortStuff(install, priorities), nil

}

func findChromeExecutables(dir string) ([]string, error) {
	_, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	chromeExecRegex := `^Exec=\/.*\/(google-chrome|chrome|chromium)-.*`
	argumentsRegex := regexp.MustCompile(`/(^[^ ]+).*/`)
	e := `grep -ER "%s" %s | awk -F '=' '{print $2}'`
	e = fmt.Sprintf(e, chromeExecRegex, dir)
	o, err := exec.Command("bash", "-c", e).Output()
	if err != nil {
		e := `grep -Er "%s" %s | awk -F '=' '{print $2}'`
		e = fmt.Sprintf(e, chromeExecRegex, dir)
		o, err = exec.Command("bash", "-c", e).Output()
		if err != nil {
			return nil, err
		}
	}
	install := []string{}
	scan := bufio.NewScanner(bytes.NewReader(o))
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		txt := scan.Text()
		txt = argumentsRegex.ReplaceAllString(txt, `$1`)
		_, err := os.Stat(txt)
		if err != nil {
			return nil, err
		}
		install = append(install, txt)
	}
	return install, nil
}
