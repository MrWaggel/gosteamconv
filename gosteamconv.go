package gosteamconv

import "strconv"

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
