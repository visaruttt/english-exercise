package main

import "english-grammar/app"

func main() {
	app := &app.GrammarTest{}
	app.Initialize()
	app.Run(":8080")
}
