package goawk

import (
	"bufio"
	"flag"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type App struct {
	Filename string
	Fs       string
	S        []string
	VS       map[string]string
	VI       map[string]int
}

type Action func(app *App)

func (app *App) Run(actions []Action) {
	errLogger := log.New(os.Stderr, "", 0)
	var fileName string

	flag.StringVar(&fileName, "i", "", "Input file name")
	flag.Parse()

	app.Filename = fileName
	app.Fs = ","
	app.VS = make(map[string]string)
	app.VI = make(map[string]int)

	firstActionName := runtime.FuncForPC(reflect.ValueOf(actions[0]).Pointer()).Name()
	if firstActionName == "main.Begin" {
		actions[0](app)
		actions = actions[1:]
	}

	lengthOfActions := len(actions)
	lastActionName := runtime.FuncForPC(reflect.ValueOf(actions[lengthOfActions-1:][0]).Pointer()).Name()
	var endAction Action

	if lastActionName == "main.End" {
		endAction = actions[lengthOfActions-1:][0]
		actions = actions[:lengthOfActions-1]
	}

	input, err := os.Open(fileName)
	defer input.Close()
	if err != nil {
		errLogger.Println("Input file does not exist.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		app.S = nil
		app.S = append(app.S, scanner.Text())
		for _, elem := range strings.Split(scanner.Text(), app.Fs) {
			app.S = append(app.S, elem)
		}
		for _, action := range actions {
			action(app)
		}
	}

	endAction(app)
}
