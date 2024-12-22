package main

import (
	"github.com/nmizern/calc_lyceum_go/internal/application"
)


func main() {
	app := application.New()
	//app.Run() запуск в консоли 
	app.RunServer()
}