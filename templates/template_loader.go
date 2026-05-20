package templates

import (
	"fmt"
	"html/template"
	"sync"

	"github.com/glebateee/core/config"
)

var once sync.Once

func LoadTemplates(cfg config.Config) error {
	path, ok := cfg.GetString("templates:path")
	if !ok {
		return fmt.Errorf("templates config path not set")
	}
	var err error
	reload := cfg.GetBoolDefault("templates:reload", true)
	once.Do(func() {
		doLoad := func() *template.Template {
			t := template.New("htmlTemplates")
			t.Funcs(template.FuncMap{
				"body":    func() string { return "" },
				"layout":  func() string { return "" },
				"handler": func() any { return "" },
			})
			t, err = t.ParseGlob(path)
			return t
		}
		if reload {
			getTemplates = doLoad
		} else {
			templates := doLoad()
			getTemplates = func() *template.Template {
				var t *template.Template
				t, err = templates.Clone()
				return t
			}
		}
	})
	return err
}
