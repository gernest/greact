package firefox

import (
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

// Pref returns default firefox preferences.
func Pref() string {
	v := []string{
		`user_pref("browser.shell.checkDefaultBrowser", false);`,
		`user_pref("browser.bookmarks.restore_default_bookmarks", false);`,
		`user_pref("dom.disable_open_during_load", false);`,
		`user_pref("dom.max_script_run_time", 0);`,
		`user_pref("dom.min_background_timeout_value", 10);`,
		`user_pref("extensions.autoDisableScopes", 0);`,
		`user_pref("browser.tabs.remote.autostart", false);`,
		`user_pref("browser.tabs.remote.autostart.2", false);`,
		`user_pref("extensions.enabledScopes", 15);`,
	}
	return strings.Join(v, "\n")
}

func getAllPrefixes() []string {
	drives := make(map[string]bool)
	p := os.Getenv("PATH")
	re := regexp.MustCompile(`^[A-Z]:\\`)
	for _, v := range strings.Split(p, ";") {
		if re.MatchString(v) && !drives[v] {
			drives[v] = true
		}
	}
	results := make(map[string]bool)
	prefixes := []string{
		os.Getenv("PROGRAMFILES"),
		os.Getenv("PROGRAMFILES(X86)"),
	}
	for _, v := range prefixes {
		for drive := range drives {
			prefix := drive + v[:1]
			if !results[prefix] {
				results[prefix] = true
			}
		}
	}
	var o []string
	for k := range results {
		o = append(o, k)
	}
	sort.Strings(o)
	return o
}

// GefFirefoxExe Return location of firefox.exe file for a given Firefox
// directory
// (available: "Mozilla Firefox", "Aurora", "Nightly").
func GefFirefoxExe(firefoxDirName string) string {
	if runtime.GOOS != "windows" {
		return ""
	}
	prefixes := getAllPrefixes()
	suffix := `\\` + firefoxDirName + `\\firefox.exe`
	for _, prefix := range prefixes {
		_, err := os.Stat(prefix + suffix)
		if err != nil {
			continue
		}
		return prefix + suffix
	}
	return `C:\\Program Files` + suffix
}
