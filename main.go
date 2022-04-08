package main

import (
	"github.com/cryptocurrency-trading-bot/controllers"
)

func main() {
	controllers.Streaming()
	controllers.StartWebServer()
}
