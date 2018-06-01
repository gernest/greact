package launcher

import (
	"os"
	"path/filepath"
)

func resolveChromePathWindows() ([]string, error) {
	var install []string
	sep := string(filepath.Separator)
	suffixes := []string{
		filepath.Join(sep, "Google", "Chrome SxS", "Application", "chrome.exe"),
		filepath.Join(sep, "Google", "Chrome", "Application", "chrome.exe"),
	}
	prefixes := []string{
		os.Getenv("LOCALAPPDATA"), os.Getenv("PROGRAMFILES"),
		os.Getenv("PROGRAMFILES(X86)"),
	}
	c := resolve()
	if c != "" {
		install = append(install, c)
	}
	for _, prefix := range prefixes {
		for _, suffix := range suffixes {
			path := filepath.Join(prefix, suffix)
			_, err := os.Stat(path)
			if err != nil {
				continue
			}
			install = append(install, path)
		}
	}
	return install, nil
}
