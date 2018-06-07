package firefox

import (
	"fmt"
	"os"
	"path/filepath"
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

// GetFirefoxExe Return location of firefox.exe file for a given Firefox
// directory
// (available: "Mozilla Firefox", "Aurora", "Nightly").
func GetFirefoxExe(firefoxDirName string) string {
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

func GetFirefoxWithFallbackOnOSX(firefoxDirNames ...string) string {
	if runtime.GOOS != "darwin" {
		return ""
	}
	prefix := "/Applications/"
	suffix := ".app/Contents/MacOS/firefox-bin"
	home := os.Getenv("HOME")
	for _, v := range firefoxDirNames {
		bin := prefix + v + suffix
		if home != "" {
			hb := filepath.Join(home, bin)
			_, err := os.Stat(hb)
			if err == nil {
				return hb
			}
		}
		_, err := os.Stat(bin)
		if err == nil {
			return bin
		}
		fmt.Println(err)
	}
	return ""
}

func firefox() string {
	switch runtime.GOOS {
	case "linux":
		return "firefox"
	case "darwin":
		return GetFirefoxWithFallbackOnOSX("Firefox")
	case "windows":
		return GetFirefoxExe("Mozilla Firefox")
	default:
		return ""
	}
}
func firefoxDeveloper() string {
	switch runtime.GOOS {
	case "linux":
		return "firefox"
	case "darwin":
		return GetFirefoxWithFallbackOnOSX("FirefoxDeveloperEdition", "FirefoxAurora")
	case "windows":
		return GetFirefoxExe("Firefox Developer Edition")
	default:
		return ""
	}
}

func firefoxAurora() string {
	switch runtime.GOOS {
	case "linux":
		return "firefox"
	case "darwin":
		return GetFirefoxWithFallbackOnOSX("FirefoxAurora")
	case "windows":
		return GetFirefoxExe("Aurora")
	default:
		return ""
	}
}
func firefoxNightly() string {
	switch runtime.GOOS {
	case "linux":
		return "firefox"
	case "darwin":
		return GetFirefoxWithFallbackOnOSX("FirefoxNightly")
	case "windows":
		return GetFirefoxExe("Nightly")
	default:
		return ""
	}
}

type Type uint

const (
	Firefox = iota
	Developer
	Aurora
	Nightly
)

func Find(typ Type) string {
	switch typ {
	case Firefox:
		return firefox()
	case Developer:
		return firefoxDeveloper()
	case Aurora:
		return firefoxAurora()
	case Nightly:
		return firefoxNightly()
	default:
		return ""
	}
}
