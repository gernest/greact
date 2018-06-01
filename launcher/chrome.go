package launcher

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

const chromePath = "CHROME_PATH"

var errChromeNotInstalled = errors.New("no Chrome installations found")

type priority struct {
	regex  *regexp.Regexp
	weight int
}
type installPriority struct {
	path   string
	weight int
}

func sortStuff(install []string, priorities []priority) []string {
	defaultPolicy := 10
	var m []*installPriority
	for _, v := range install {
		for _, p := range priorities {
			if p.regex.MatchString(v) {
				m = append(m, &installPriority{
					path: v, weight: p.weight,
				})
				continue
			}
		}
		m = append(m, &installPriority{
			path: v, weight: defaultPolicy,
		})
	}
	sort.Slice(m, func(a, b int) bool {
		return m[a].weight < m[b].weight
	})
	var o []string
	for _, v := range m {
		o = append(o, v.path)
	}
	return o
}

func defaultFlags() []string {
	return []string{
		"--disable-translate",
		"--disable-extensions",
		"--disable-background-networking",
		"--safebrowsing-disable-auto-update",
		"--disable-sync",
		"--metrics-recording-only",
		"--disable-default-apps",
		"--mute-audio",
		"--no-first-run",
	}
}

type Options struct {
	StartingURL        string
	ChromeFlags        []string
	Port               int
	ChromePath         string
	IgnoreDefaultFlags bool
	UserDataDir        string

	//The time taken to wait for chrome to be ready.
	WaitTimeout time.Duration
}

func (o Options) Flags() []string {
	var f []string
	if !o.IgnoreDefaultFlags {
		f = defaultFlags()
	}
	f = append(f, "--remote-debugging-port="+fmt.Sprint(o.Port))
	if runtime.GOOS == "linux" {
		f = append(f, "--disable-setuid-sandbox")
	}
	f = append(f, o.ChromeFlags...)
	f = append(f, o.StartingURL)
	return f
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
	if opts.ChromeFlags == nil {
		opts.ChromeFlags = []string{}
	}
	if opts.UserDataDir == "" {
		dir := filepath.Join(os.TempDir(), "mad-launcher")
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
		opts.UserDataDir = dir
	}
	if opts.ChromePath == "" {
		i, err := resolveChromePath()
		if err != nil {
			return nil, err
		}
		if len(i) == 0 {
			return nil, errChromeNotInstalled
		}
		opts.ChromePath = i[0]
	}
	if opts.Port == 0 {
		v, err := randomPort()
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		opts.Port = i
	}
	if opts.WaitTimeout == 0 {
		opts.WaitTimeout = 2 * time.Second
	}
	l := &Launcher{Opts: opts}
	ctx, cancel := context.WithCancel(context.Background())
	l.Cmd = exec.CommandContext(ctx, l.Opts.ChromePath, l.Opts.Flags()...)
	l.cancel = cancel
	return l, nil
}

func getPlatform() (string, error) {
	v := runtime.GOOS
	switch v {
	case "darwin", "linux":
		return v, nil
	default:
		return "", fmt.Errorf("platform %s is not supported", v)
	}
}

func resolveChromePath() ([]string, error) {
	platform, err := getPlatform()
	if err != nil {
		return nil, err
	}
	switch platform {
	case "darwin":
		return resolveChromePathDarwin()
	case "linux":
		return resolveChromePathLinux()
	default:
		return nil, fmt.Errorf("platform %s is not supported", platform)
	}
}

func (l *Launcher) Run() error {
	return l.Cmd.Run()
}

func (l *Launcher) Stop() {
	l.cancel()
}

func (l *Launcher) Wait(verbose bool) error {
	if verbose {
		fmt.Print("Waiting for chrome ...")
	}
	status := "."
	tick := time.NewTicker(time.Second / 3)
	defer tick.Stop()
	o := time.NewTimer(l.Opts.WaitTimeout)
	defer o.Stop()
	for {
		select {
		case <-o.C:
			if verbose {
				fmt.Println(".")
			}
			return errors.New("timeout waiting for chrome to be ready")
		case <-tick.C:
			conn, err := net.Dial("tcp", fmt.Sprintf(":%d", l.Opts.Port))
			if err != nil {
				if verbose {
					fmt.Print(status)
				}
				status += "."
				continue
			}
			conn.Close()
			if verbose {
				fmt.Println("done")
			}
			return nil
		}
	}
}

func randomPort() (string, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	err = l.Close()
	if err != nil {
		return "", err
	}
	a := l.Addr().String()
	println(a)
	if strings.HasSuffix(a, ":") {
		a = a[:len(a)-1]
	}
	p := strings.Split(a, ":")
	return p[len(p)-1], nil
}
