package main

import "github.com/tawesoft/golib/v2/dialog"

func main() {
	err := dialog.Info("Hello World")
	if err != nil {
		return
	}
}
