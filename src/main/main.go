package main

import (
	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

var version string

func main() {
	result := app.New(version).Run()

	utils.OsExit(result)
}
