package main

import (
	"fmt"
	"github.com/HordeGroup/horde/pkg/webserver"
	"os"
)

func GetServerRunnerFunc(appMode string) func() error {
	switch appMode {
	case "web":
		return webserver.Run
	default:
		panic(fmt.Sprintf("unknown app mode %s", appMode))
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("not enough arguments")
	}

	appMode := os.Args[1]
	runnerFn := GetServerRunnerFunc(appMode)
	if err := runnerFn(); err != nil {
		panic(err)
	}
}
