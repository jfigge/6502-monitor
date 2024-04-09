/*
 * Copyright (C) 2024 by Jason Figge
 */

package internal

import (
	"us.figge.monitor/internal/configuration"
)

type Application struct {
	cm *configuration.ConfigManager
}

func NewApplication() (*Application, bool) {
	app := &Application{}
	var ok bool
	if app.cm, ok = configuration.NewConfigManager(); !ok {
		return nil, false
	}
	return app, true
}

func (a *Application) Configuration() *configuration.ConfigManager {
	return a.cm
}

func (a *Application) Start() {

}
