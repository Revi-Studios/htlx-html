package main

import (
	"htlx/cmd/parser"
	"syscall/js"
)

func main() {
	js.Global().Set("parse", js.FuncOf(func(this js.Value, args []js.Value) any {
        if len(args) < 1 { return nil }
        input := args[0].String()
        return parser.Parse(input)
	}
	)))
	select {}
}