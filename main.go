package main

import (
	"log"
	"ticket/app"
)

func main() {
	app := app.NewApp()
	log.Fatalln(app.Start(":" + "3000")) // todo change port based on configs
}
