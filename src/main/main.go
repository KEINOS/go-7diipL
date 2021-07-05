package main

import (
	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

func main() {
	result := app.New().Run()

	utils.OsExit(result)
}
