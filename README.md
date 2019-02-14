# gosteamconv
This package is meant to convert the steamid format 'STEAM_0:0:123456' to an 32 or 64 bit integer. This is all done following the official wiki.

https://developer.valvesoftware.com/wiki/SteamID

## go get
```
go get github.com/MrWaggel/gosteamconv/
```

## GoDoc
https://godoc.org/github.com/MrWaggel/gosteamconv

[![GoDoc](https://godoc.org/github.com/MrWaggel/gosteamconv?status.svg)](https://godoc.org/github.com/MrWaggel/gosteamconv)

## Example
```go
package main

import (
	"fmt"
	"github.com/MrWaggel/gosteamconv"
)

func main() {
	steam64, err := gosteamconv.SteamStringToInt64("STEAM_0:0:6545518")
	if err == nil {
		fmt.Println(steam64)
	} else {
		fmt.Println(err)
	}
}
```
Will output the following
```
76561197973356764
```
http://steamcommunity.com/profiles/76561197973356764
