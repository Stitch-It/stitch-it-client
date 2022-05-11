//go:build wasm
// +build wasm

package main

import (
	gen "github.com/Stitch-It/stitch-it/go/generate-pattern"
	imgHdl "github.com/Stitch-It/stitch-it/go/image-process"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("processAndCreatePattern", js.FuncOf(processAndCreatePattern))
	<-done
}

// May need to change into expecting a bas64 encoded string
// not sure because syscall/js doesn't provide a type cast to
// a []byte, but does provide a type cast to String
func processAndCreatePattern(this js.Value, args []js.Value) interface{} {
	resImg := imgHdl.ResizeImage(args[0].String(), args[1].Bool(), args[3].Int(), args[4].Int())

	buf := gen.GenerateExcelPattern(resImg)

	return buf
}
