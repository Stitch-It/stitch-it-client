Instructions to come!

### Note for cross compiling to WASM on Windows ###

On Windows, to cross compile to WASM, instead of simply running

```
GOOS=js GOARCH=wasm go build -o static/main.wasm main.go
```

you must run

```
$Env:GOOS = "js"; $Env:GOARCH = "wasm"; go build -o static/main.wasm main.go"
```

Once the full website is built out, this will become the client repo as the WASM module must be available for the client to use. The rest of the Go files (those for routing, processing payments via the Go Stripe SDK, etc.) will be back-end, and likely reside in their own repo
