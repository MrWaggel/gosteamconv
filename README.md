# gosteamconv
This package is meant to convert the steamid format 'STEAM_0:0:123456' to an 32 or 64 bit integer. This is all done following the official wiki.

https://developer.valvesoftware.com/wiki/SteamID

##Install
```
go get github.com/MrWaggel/gosteamconv/
```

##Functions
```
SteamStringToInt32(string) (int, error)
```
Converts the STEAM_ string to an 32bit integer.

```
SteamStringToInt64(string) (int64, error)
```
Converts the STEAM_ string into an 64bit integer (community id)

###Example
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
	}
	fmt.Println(steam64)
}
```
Will output the following
```
76561197973356764
```
http://steamcommunity.com/profiles/76561197973356764
