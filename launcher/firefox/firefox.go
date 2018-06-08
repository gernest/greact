package firefox

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"
)

var errFirefoxNotInstalled = errors.New("no Firefox installations found")

const DefaultPort = 6000

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

type Options struct {
	StartingURL        string
	FirefoxFlags       []string
	Type               Type
	Port               int
	FirefoxPath        string
	IgnoreDefaultFlags bool
	UserDataDir        string

	//The time taken to wait for chrome to be ready.
	WaitTimeout time.Duration

	ProfilePath string
	Verbose     bool
}

func (o Options) Flags() []string {
	var f []string
	if !o.IgnoreDefaultFlags {
		f = defaultFlags()
	}
	f = append(f, "-start-debugger-server="+fmt.Sprint(o.Port),
		"-profile="+o.ProfilePath,
	)
	f = append(f, o.FirefoxFlags...)
	f = append(f, o.StartingURL)
	return f
}

func defaultFlags() []string {
	return []string{}
}

type Launcher struct {
	Opts   Options
	Cmd    *exec.Cmd
	cancel func()
}

func New(opts Options) (*Launcher, error) {
	if opts.StartingURL == "" {
		opts.StartingURL = "about:blank"
	}
	if opts.FirefoxFlags == nil {
		opts.FirefoxFlags = []string{}
	}
	if opts.UserDataDir == "" {
		dir := filepath.Join(os.TempDir(), "mad-launcher")
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
		p := filepath.Join(dir, "firefox_profile.js")
		err = ioutil.WriteFile(p, []byte(Pref()), 0600)
		if err != nil {
			return nil, err
		}
		opts.ProfilePath = p
		opts.UserDataDir = dir
	}
	if opts.FirefoxPath == "" {
		i := Find(opts.Type)
		if i == "" {
			return nil, errFirefoxNotInstalled
		}
		opts.FirefoxPath = i
	}
	if opts.Port == 0 {
		opts.Port = DefaultPort
	}
	if opts.WaitTimeout == 0 {
		opts.WaitTimeout = 30 * time.Second
	}
	return &Launcher{Opts: opts}, nil
}

func (l *Launcher) Run(bctx context.Context) error {
	ctx, cancel := context.WithCancel(bctx)
	l.Cmd = exec.CommandContext(ctx, l.Opts.FirefoxPath, l.Opts.Flags()...)
	l.Cmd.Stdout = os.Stdout
	l.Cmd.Stderr = os.Stdout
	fmt.Println(l.Cmd.Args)
	l.cancel = cancel
	return l.Cmd.Run()
}

func (l *Launcher) Stop() error {
	l.cancel()
	return nil
}

func (l *Launcher) Ready() error {
	if l.Opts.Verbose {
		fmt.Print("Waiting for firefox ...")
	}
	status := "."
	tick := time.NewTicker(time.Second / 3)
	defer tick.Stop()
	o := time.NewTimer(l.Opts.WaitTimeout)
	defer o.Stop()
	for {
		select {
		case <-o.C:
			if l.Opts.Verbose {
				fmt.Println(".")
			}
			return errors.New("timeout waiting for firefox to be ready")
		case <-tick.C:
			conn, err := net.Dial("tcp", fmt.Sprintf(":%d", l.Opts.Port))
			if err != nil {
				if l.Opts.Verbose {
					fmt.Print(status)
				}
				status += "."
				continue
			}
			conn.Close()
			if l.Opts.Verbose {
				fmt.Println("done")
			}
			return nil
		}
	}
}
