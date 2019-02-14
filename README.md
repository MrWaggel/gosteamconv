# gosteamconv [![GoDoc](https://godoc.org/github.com/MrWaggel/gosteamconv?status.svg)](https://godoc.org/github.com/MrWaggel/gosteamconv)
This package is meant to convert the steamid format 'STEAM_0:0:123456' to a 32 or 64 bit integer, or an integer to the string format. This is all done following the official wiki.
https://developer.valvesoftware.com/wiki/SteamID

## go get
```
go get github.com/MrWaggel/gosteamconv
```

## Example
```go
package main

import (
	"fmt"
	"github.com/MrWaggel/gosteamconv"
)

func main() {
	steam64, err := gosteamconv.SteamStringToInt64("STEAM_0:0:6545518")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(steam64)
	}
}
```
Will output the following
```
76561197973356764
```
