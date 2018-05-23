package launcher

// func resolveChromePathLinux() ([]string, error) {
// 	var install []string
// 	c := resolve()
// 	if c != "" {
// 		install = append(install, c)
// 	}
// 	u, err := user.Current()
// 	if err != nil {
// 		return nil, err
// 	}
// 	desktonInstall := []string{
// 		filepath.Join(u.HomeDir, ".local/share/applications/"),
// 		"/usr/share/applications",
// 	}

// }

// func findChromeExecutables(dir string) ([]string, error) {
// 	_, err := os.Stat(dir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	chromeExecRegex := `^Exec=\/.*\/(google-chrome|chrome|chromium)-.*`
// 	argumentsRegex := `/(^[^ ]+).*/`
// 	e := `grep -ER "%s" %s | awk -F '=' '{print $2}'`
// 	e = fmt.Sprintf(e, chromeExecRegex, dir)
// 	o, err := exec.Command("bash", "-c", e).Output()
// 	if err != nil {
// 		e := `grep -Er "%s" %s | awk -F '=' '{print $2}'`
// 		e = fmt.Sprintf(e, chromeExecRegex, dir)
// 		o, err = exec.Command("bash", "-c", e).Output()
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// }
