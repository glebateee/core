package services

import (
	"github.com/glebateee/core/config"
	"github.com/glebateee/core/logging"
	"github.com/glebateee/core/templates"
	"github.com/glebateee/core/validation"
)

func RegisterDefaultServices() {
	if err := AddSingleton(func() config.Config {
		cfg, err := config.Load("config.json")
		if err != nil {
			panic(err)
		}
		return cfg
	}); err != nil {
		panic(err)
	}
	if err := AddSingleton(func(cfg config.Config) logging.Logger {
		return logging.NewDefaultLogger(cfg)
	}); err != nil {
		panic(err)
	}

	if err := AddSingleton(func(cfg config.Config) templates.TemplateExecutor {
		if err := templates.LoadTemplates(cfg); err != nil {
			panic(err)
		}
		return &templates.LayoutTemplateProcessor{}
	}); err != nil {
		panic(err)
	}
	if err := AddSingleton(func() validation.Validator {
		return validation.NewTagValidator(validation.DefaultValidators())
	}); err != nil {
		panic(err)
	}
}
