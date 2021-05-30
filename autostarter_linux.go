// +build linux

package autostarter

import (
	"errors"
	"os"
	"path/filepath"
	"text/template"
)

const (
	shortcutTemplate = `[Desktop Entry]
Name={{.DisplayName}}
Exec={{.Exec}}
Path={{.Path}}
{{- if .Icon}}
Icon={{.Icon}}{{end}}
Type=Application
X-GNOME-Autostart-enabled=true
`
	shortcutExt = ".desktop"
)

type shortcut struct {
	DisplayName string
	Exec        string
	Path        string
	Icon        icon
}

func getStartupDir() string {
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		return os.Getenv("XDG_CONFIG_HOME")
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "autostart")
}

func checkIcon(path string) error {
	switch filepath.Ext(path) {
	case ".ico":
	case ".png":
	case ".jpg":
	default:
		return errors.New("format invalid: " + filepath.Ext(path))
	}
	return nil
}

func createShortcut(sc Shortcut, ic icon) (err error) {
	var exec string
	t := template.Must(template.New("shortcut").Parse(shortcutTemplate))
	f, err := os.Create(filepath.Join(getStartupDir(), sc.Name+shortcutExt))
	if err != nil {
		return
	}
	defer f.Close()
	if len(sc.Args) > 0 {
		exec = sc.Exec + " " + sc.getArgsString()
	} else {
		exec = sc.Exec
	}
	t.Execute(f, &shortcut{
		DisplayName: sc.Name,
		Exec:        exec,
		Path:        sc.StartIn,
		Icon:        ic,
	})
	return
}
