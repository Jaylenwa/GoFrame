package cmd

import (
	"GoFrame/application/adapter/conncfg"
)

type App struct {
	ConncfgSrv *adapter.ConncfgAdapter
}

func InitApp() *App {
	return &App{
		ConncfgSrv: adapter.NewConncfgAdapter(),
	}
}
