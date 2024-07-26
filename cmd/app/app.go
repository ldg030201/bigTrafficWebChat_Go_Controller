package app

import "chat_controller_server/config"

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	a := &App{cfg: cfg}

	return a
}
