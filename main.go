package main

import (
	"fmt"
	"interview-accountapi/app"
	"github.com/fatih/color"
	"rsc.io/quote"
)

func main() {
	fmt.Println(app.Env("HOME"))
	color.Cyan(quote.Hello())
}
