package gosteamconv

import "testing"

func BenchmarkSteamInt32ToString(b *testing.B) {
	id := int32(13091036)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SteamInt32ToString(id)
	}
}

func BenchmarkSteamInt64ToString(b *testing.B) {
	id := int64(76561197973356764)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SteamInt64ToString(id)
	}
}

func BenchmarkSteamStringToInt32(b *testing.B) {
	str := "STEAM_0:0:6545518"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SteamStringToInt32(str)
	}
}

func BenchmarkSteamStringToInt64(b *testing.B) {
	str := "STEAM_0:0:6545518"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SteamStringToInt64(str)
	}
}
