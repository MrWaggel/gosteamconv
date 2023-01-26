package gosteamconv

import "testing"

func TestSteamInt32ToString(t *testing.T) {
	i := int32(13091036)

	str, err := SteamInt32ToString(i)
	if err != nil {
		t.Error(err)
		return
	}

	if str != "STEAM_0:0:6545518" {
		t.Error("conversion incorrect")
	}
}

func TestSteamInt64ToString(t *testing.T) {
	i := int64(76561197973356764)

	str, err := SteamInt64ToString(i)
	if err != nil {
		t.Error(err)
		return
	}

	if str != "STEAM_0:0:6545518" {
		t.Error("conversion incorrect")
	}
}

func TestSteamStringToInt32(t *testing.T) {
	s := "STEAM_0:0:6545518"

	i, err := SteamStringToInt32(s)
	if err != nil {
		t.Error(err)
		return
	}

	if i != 13091036 {
		t.Error("conversion incorrect")
	}
}

func TestSteamStringToInt64(t *testing.T) {
	s := "STEAM_0:0:6545518"

	i, err := SteamStringToInt64(s)
	if err != nil {
		t.Error(err)
		return
	}

	if i != 76561197973356764 {
		t.Error("conversion incorrect")
	}
}
