package main

import (
	"github.com/goravel/framework/facades"
	"fmt"
	"goravel/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	//Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
		fmt.Println("hello world")
	}()

	select {}
}
