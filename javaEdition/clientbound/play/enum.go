package play

type PlayerInfoMask = uint8

const (
	PlayerInfoMaskAddPlayer         PlayerInfoMask = 0x01
	PlayerInfoMaskInitializeChannel PlayerInfoMask = 0x02
	PlayerInfoMaskUpdateGameMode    PlayerInfoMask = 0x04
	PlayerInfoMaskUpdateListed      PlayerInfoMask = 0x08
	PlayerInfoMaskUpdateLatency     PlayerInfoMask = 0x10
	PlayerInfoMaskUpdateDisplayName PlayerInfoMask = 0x20
)
