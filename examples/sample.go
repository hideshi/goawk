package main

import (
	"github.com/hideshi/goawk"
	"fmt"
    "strconv"
    "regexp"
)

func Begin(app *goawk.App) {
	fmt.Println("Beginを実行します。")
	fmt.Println(app.Filename)
}

func Action1(app *goawk.App) {
	fmt.Println("Action1を実行します。")
	fmt.Println("入力は１行ごとに処理され、文字列はフィールドセパレータで分割されます。")
    for _, elem := range app.S[1:] {
	    fmt.Printf("%#v\n", elem)
        v, err := strconv.Atoi(elem)
        if err != nil {
            fmt.Println(err)
        } else {
            app.VI["sum"] = app.VI["sum"] + v
        }
    }
}

func Action2(app *goawk.App) {
	fmt.Println("Action2を実行します。")
    pattern := ".*20.*"
    matched, _ := regexp.MatchString(pattern, app.S[0])
    if matched == true {
        fmt.Printf("%sが%sのパターンにマッチしました。\n", app.S[0], pattern)
    }
}

func End(app *goawk.App) {
	fmt.Println("Endを実行します。")
	fmt.Printf("合計は%dです。\n", app.VI["sum"])
}

func main() {
	app := new(goawk.App)
	actions := []goawk.Action{Begin, Action1, Action2, End}
	app.Run(actions)
}
