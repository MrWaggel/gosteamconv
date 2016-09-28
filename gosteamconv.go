// Package gosteamconv provides methods to convert steamids between strings (STEAM_X:Y:Z) and integers like the Steam community id.
package gosteamconv

import (
	"errors"
	"strconv"
)

// SteamStringToInt32 takes a steamid string "STEAM_X:Y:Z" and converts it to a 32 bit integer.
func SteamStringToInt32(steamstring string) (int, error) {
	Y, err := strconv.Atoi(steamstring[8:9])
	if err != nil {
		return int(0), err
	}
	Z, err := strconv.Atoi(steamstring[10:])
	if err != nil {
		return int(0), err
	}
	return (Z * 2) + Y, nil
}

// SteamStringToInt64 takes a steamid string "STEAM_X:Y:Z" and converts it to a 64 bit integer.
func SteamStringToInt64(steamstring string) (int64, error) {
	Y, err := strconv.Atoi(steamstring[8:9])
	if err != nil {
		return int64(0), err
	}
	Z, err := strconv.Atoi(steamstring[10:])
	if err != nil {
		return int64(0), err
	}
	return int64((Z * 2) + 76561197960265728 + Y), nil
}

// SteamInt64ToString takes a 64 bit integer and converts it as a steamid as a steamid string format "STEAM_X:Y:Z"
// The argument must be bigger than 76561197960265728 or it will return an error.
func SteamInt64ToString(steamInt int64) (string, error) {
	if steamInt <= 76561197960265728 {
		return string(""), errors.New("64 bit steamid int should be bigger than 76561197960265728")
	}
	steamInt = steamInt - 76561197960265728
	remainder := steamInt % 2
	steamInt = steamInt / 2
	return "STEAM_0:" + strconv.FormatInt(remainder, 10) + ":" + strconv.FormatInt(steamInt, 10), nil

}

// SteamInt64ToString takes a 32 bit integer and converts it as a steamid as a steamid string format "STEAM_X:Y:Z"
// The argument must be bigger than 0 or it will return an error.
func SteamInt32ToString(steamInt int32) (string, error) {
	if steamInt <= 0 {
		return string(""), errors.New("32 bit steamid int should be bigger than 0")
	}
	remainder := steamInt % 2
	steamInt = steamInt / 2
	return "STEAM_0:" + strconv.FormatInt(int64(remainder), 10) + ":" + strconv.FormatInt(int64(steamInt), 10), nil
}
