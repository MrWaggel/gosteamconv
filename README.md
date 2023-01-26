# gosteamconv
[![Go Reference](https://pkg.go.dev/static/frontend/badge/badge.svg)](https://pkg.go.dev/github.com/mrwaggel/gosteamconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrwaggel/gosteamconv)](https://goreportcard.com/report/github.com/mrwaggel/gosteamconv)

This Go package provides functions to convert the steamid string format 'STEAM_0:0:123456' to a 32 or 64 bit integer, 
or an integer to the string format. This is all done following the official wiki.
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
