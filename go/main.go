//go:build wasm
// +build wasm

package main

import (
	"syscall/js"

	gen "github.com/Stitch-It/stitch-it/go/generate-pattern"
	imgHdl "github.com/Stitch-It/stitch-it/go/image-process"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("processAndCreatePattern", js.FuncOf(processAndCreatePattern))
	<-done
}

func processAndCreatePattern(this js.Value, args []js.Value) interface{} {
	resImg := imgHdl.ResizeImage(args[0].String(), args[1].Bool(), args[3].Int(), args[4].Int())

	buf := gen.GenerateExcelPattern(resImg)

	return buf
}
