package main

import (
	"fmt"
	"github.com/hideshi/goawk"
)

func Action1(app *goawk.App) {
	app.VI[app.S[0]] = app.VI[app.S[0]] + 1
}

func End(app *goawk.App) {
    for k, v := range(app.VI) {
	    fmt.Printf("%v : %v\n", k, v)
    }
}

func main() {
	app := new(goawk.App)
	actions := []goawk.Action{Action1, End}
	app.Run(actions)
}
