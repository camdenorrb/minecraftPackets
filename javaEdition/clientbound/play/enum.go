package play

type PlayerInfoMask = uint8

const (
	PlayerInfoAddPlayer         PlayerInfoMask = 0x01
	PlayerInfoInitializeChannel PlayerInfoMask = 0x02
	PlayerInfoUpdateGameMode    PlayerInfoMask = 0x04
	PlayerInfoUpdateListed      PlayerInfoMask = 0x08
	PlayerInfoUpdateLatency     PlayerInfoMask = 0x10
	PlayerInfoUpdateDisplayName PlayerInfoMask = 0x20
)
