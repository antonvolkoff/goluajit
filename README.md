# LuaJit binding to Go

## Installing

`go get -u github.com/choix/goluajit`

## Examples

```go
package main

import (
	"fmt"
	"github.com/choix/goluajit"
)

func main() {
	L := luajit.NewState()
	defer L.Close()
	L.OpenLibs()

	L.LoadString(`print("Hello World!")`)
	L.Run()
}
```
