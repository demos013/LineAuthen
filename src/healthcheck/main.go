package main

import (
	"healthcheck/config"
	"healthcheck/router"
	"healthcheck/task"
	"os"
)

func main() {

	if len(os.Args) >= 2 {
		task.CheckWebsite(os.Args[len(os.Args)-1])
		return
	}

	config := config.New()
	app := router.New(config)
	app.Run(":" + config.AppPort())

}
