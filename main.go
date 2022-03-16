package main

//To get All echo framework we have to run this command
//go get -u github.com/labstack/echo/...       //u means update
//go get github.com/labstack/echo/v4

import (
	"echoweb/tronics"
)

func main() {
	tronics.Text()
	tronics.Start()
}
