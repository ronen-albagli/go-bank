package main

import (
	"bank/webserver"
	"fmt"
)

func main() {
	app := webserver.CreateWebServer()

	app.Run("localhost:8002")
	fmt.Println("Hello world")
}
