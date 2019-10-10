package main

import (
	"os"

	"github.com/theantichris/echo"
)

func main() {
	echo.Echo(os.Args, os.Stdout)
}
