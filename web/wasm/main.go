//go:build js && wasm

package main

import "syscall/js"

func main() {
	c := make(chan struct{})
	js.Global().Set("helloFromWasm", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return "hello from wasm"
	}))
	<-c
}
