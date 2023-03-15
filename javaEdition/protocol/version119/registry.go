package version119

import (
	clientLogin "github.com/camdenorrb/minecraftPackets/javaEdition/clientbound/login"
	clientPlay "github.com/camdenorrb/minecraftPackets/javaEdition/clientbound/play"
	clientStatus "github.com/camdenorrb/minecraftPackets/javaEdition/clientbound/status"
	"github.com/camdenorrb/minecraftPackets/javaEdition/protocol"
	serverHandshake "github.com/camdenorrb/minecraftPackets/javaEdition/serverbound/handshake"
	serverLogin "github.com/camdenorrb/minecraftPackets/javaEdition/serverbound/login"
	serverPlay "github.com/camdenorrb/minecraftPackets/javaEdition/serverbound/play"
	serverStatus "github.com/camdenorrb/minecraftPackets/javaEdition/serverbound/status"
)

func Registry(bound protocol.Bound) protocol.PacketRegistry {

	switch bound {
	case protocol.ClientBound:
		return clientBound()
	case protocol.ServerBound:
		return serverBound()
	}

	return nil
}

func clientBound() protocol.PacketRegistry {

	registry := protocol.NewPacketRegistry()

	// Status
	registry.Register(protocol.StatusState, 0x00, clientStatus.Response{})
	registry.Register(protocol.StatusState, 0x01, clientStatus.PingResponse{})

	// Login
	registry.Register(protocol.LoginState, 0x00, clientLogin.Disconnect{})
	registry.Register(protocol.LoginState, 0x01, clientLogin.EncryptionRequest{})
	registry.Register(protocol.LoginState, 0x02, clientLogin.Success{})
	registry.Register(protocol.LoginState, 0x03, clientLogin.SetCompression{})
	registry.Register(protocol.LoginState, 0x04, clientLogin.PluginRequest{})

	// Play
	registry.Register(protocol.PlayState, 0x00, clientPlay.SpawnEntity{})
	registry.Register(protocol.PlayState, 0x01, clientPlay.SpawnExperienceOrb{})
	registry.Register(protocol.PlayState, 0x02, clientPlay.SpawnPlayer{})
	registry.Register(protocol.PlayState, 0x03, clientPlay.EntityAnimation{})
	registry.Register(protocol.PlayState, 0x04, clientPlay.AwardStatistics{})
	registry.Register(protocol.PlayState, 0x05, clientPlay.AcknowledgeBlockChange{})
	registry.Register(protocol.PlayState, 0x06, clientPlay.SetBlockDestroyStage{})
	registry.Register(protocol.PlayState, 0x07, clientPlay.BlockEntityData{})
	registry.Register(protocol.PlayState, 0x08, clientPlay.BlockAction{})
	registry.Register(protocol.PlayState, 0x09, clientPlay.BlockUpdate{})
	registry.Register(protocol.PlayState, 0x0A, clientPlay.BossBar{})
	registry.Register(protocol.PlayState, 0x0B, clientPlay.ChangeDifficulty{})
	registry.Register(protocol.PlayState, 0x0C, clientPlay.ClearTitles{})
	registry.Register(protocol.PlayState, 0x0D, clientPlay.CommandSuggestionsResponse{})
	registry.Register(protocol.PlayState, 0x0E, clientPlay.Commands{})
	registry.Register(protocol.PlayState, 0x0F, clientPlay.CloseContainer{})
	registry.Register(protocol.PlayState, 0x10, clientPlay.SetContainerContent{})
	registry.Register(protocol.PlayState, 0x11, clientPlay.SetContainerProperty{})
	registry.Register(protocol.PlayState, 0x12, clientPlay.SetContainerSlot{})
	registry.Register(protocol.PlayState, 0x13, clientPlay.SetCooldown{})
	registry.Register(protocol.PlayState, 0x14, clientPlay.ChatSuggestions{})
	registry.Register(protocol.PlayState, 0x15, clientPlay.PluginMessage{})
	registry.Register(protocol.PlayState, 0x16, clientPlay.DeleteMessage{})
	registry.Register(protocol.PlayState, 0x17, clientPlay.Disconnect{})
	registry.Register(protocol.PlayState, 0x18, clientPlay.DisguisedChatMessage{})
	registry.Register(protocol.PlayState, 0x19, clientPlay.EntityEvent{})
	registry.Register(protocol.PlayState, 0x1A, clientPlay.Explosion{})
	registry.Register(protocol.PlayState, 0x1B, clientPlay.UnloadChunk{})
	registry.Register(protocol.PlayState, 0x1C, clientPlay.GameEvent{})
	registry.Register(protocol.PlayState, 0x1D, clientPlay.OpenHorseScreen{})
	registry.Register(protocol.PlayState, 0x1E, clientPlay.InitializeWorldBorder{})
	registry.Register(protocol.PlayState, 0x1F, clientPlay.KeepAlive{})
	registry.Register(protocol.PlayState, 0x20, clientPlay.ChunkDataAndUpdateLight{})
	registry.Register(protocol.PlayState, 0x21, clientPlay.WorldEvent{})
	registry.Register(protocol.PlayState, 0x22, clientPlay.Particle{})
	registry.Register(protocol.PlayState, 0x23, clientPlay.UpdateLight{})
	registry.Register(protocol.PlayState, 0x24, clientPlay.Login{})
	registry.Register(protocol.PlayState, 0x25, clientPlay.MapData{})
	registry.Register(protocol.PlayState, 0x26, clientPlay.MerchantOffers{})
	registry.Register(protocol.PlayState, 0x27, clientPlay.UpdateEntityPosition{})
	registry.Register(protocol.PlayState, 0x28, clientPlay.UpdateEntityPositionAndRotation{})
	registry.Register(protocol.PlayState, 0x29, clientPlay.UpdateEntityRotation{})
	registry.Register(protocol.PlayState, 0x2A, clientPlay.MoveVehicle{})
	registry.Register(protocol.PlayState, 0x2B, clientPlay.OpenBook{})
	registry.Register(protocol.PlayState, 0x2C, clientPlay.OpenWindow{})
	registry.Register(protocol.PlayState, 0x2D, clientPlay.OpenSignEditor{})
	registry.Register(protocol.PlayState, 0x2E, clientPlay.Ping{})
	registry.Register(protocol.PlayState, 0x2F, clientPlay.PlaceGhostRecipe{})
	registry.Register(protocol.PlayState, 0x30, clientPlay.PlayerAbilities{})
	registry.Register(protocol.PlayState, 0x31, clientPlay.PlayerChatMessage{})
	registry.Register(protocol.PlayState, 0x32, clientPlay.EndCombat{})
	registry.Register(protocol.PlayState, 0x33, clientPlay.EnterCombat{})
	registry.Register(protocol.PlayState, 0x34, clientPlay.CombatDeath{})
	registry.Register(protocol.PlayState, 0x35, clientPlay.PlayerInfoRemove{})
	registry.Register(protocol.PlayState, 0x36, clientPlay.PlayerInfoUpdate{})
	registry.Register(protocol.PlayState, 0x37, clientPlay.LookAt{})
	registry.Register(protocol.PlayState, 0x38, clientPlay.SynchronizePlayerPosition{})
	registry.Register(protocol.PlayState, 0x39, clientPlay.UpdateRecipeBook{})
	registry.Register(protocol.PlayState, 0x3A, clientPlay.RemoveEntities{})
	registry.Register(protocol.PlayState, 0x3B, clientPlay.RemoveEntityEffect{})
	registry.Register(protocol.PlayState, 0x3C, clientPlay.ResourcePack{})
	registry.Register(protocol.PlayState, 0x3D, clientPlay.Respawn{})
	registry.Register(protocol.PlayState, 0x3E, clientPlay.SetHeadRotation{})
	registry.Register(protocol.PlayState, 0x3F, clientPlay.UpdateSectionBlocks{})
	registry.Register(protocol.PlayState, 0x40, clientPlay.SelectAdvancementTab{})
	registry.Register(protocol.PlayState, 0x41, clientPlay.ServerData{})
	registry.Register(protocol.PlayState, 0x42, clientPlay.SetActionBarText{})
	registry.Register(protocol.PlayState, 0x43, clientPlay.SetBorderCenter{})
	registry.Register(protocol.PlayState, 0x44, clientPlay.SetBorderLerpSize{})
	registry.Register(protocol.PlayState, 0x45, clientPlay.SetBorderSize{})
	registry.Register(protocol.PlayState, 0x46, clientPlay.SetBorderWarningDelay{})
	registry.Register(protocol.PlayState, 0x47, clientPlay.SetBorderWarningDistance{})
	registry.Register(protocol.PlayState, 0x48, clientPlay.SetCamera{})
	registry.Register(protocol.PlayState, 0x49, clientPlay.SetHeldItem{})
	registry.Register(protocol.PlayState, 0x4A, clientPlay.SetCenterChunk{})
	registry.Register(protocol.PlayState, 0x4B, clientPlay.SetRenderDistance{})
	registry.Register(protocol.PlayState, 0x4C, clientPlay.SetDefaultSpawnPosition{})
	registry.Register(protocol.PlayState, 0x4D, clientPlay.DisplayObjective{})
	registry.Register(protocol.PlayState, 0x4E, clientPlay.SetEntityMetadata{})
	registry.Register(protocol.PlayState, 0x4F, clientPlay.LinkEntities{})
	registry.Register(protocol.PlayState, 0x50, clientPlay.SetEntityVelocity{})
	registry.Register(protocol.PlayState, 0x51, clientPlay.SetEquipment{})
	registry.Register(protocol.PlayState, 0x52, clientPlay.SetExperience{})
	registry.Register(protocol.PlayState, 0x53, clientPlay.SetHealth{})
	registry.Register(protocol.PlayState, 0x54, clientPlay.UpdateObjectives{})
	registry.Register(protocol.PlayState, 0x55, clientPlay.SetPassengers{})
	registry.Register(protocol.PlayState, 0x56, clientPlay.UpdateTeams{})
	registry.Register(protocol.PlayState, 0x57, clientPlay.UpdateScore{})
	registry.Register(protocol.PlayState, 0x58, clientPlay.SetSimulationDistance{})
	registry.Register(protocol.PlayState, 0x59, clientPlay.SetSubtitleText{})
	registry.Register(protocol.PlayState, 0x5A, clientPlay.UpdateTime{})
	registry.Register(protocol.PlayState, 0x5B, clientPlay.SetTitleText{})
	registry.Register(protocol.PlayState, 0x5C, clientPlay.SetTitleAnimationTimes{})
	registry.Register(protocol.PlayState, 0x5D, clientPlay.EntitySoundEffect{})
	registry.Register(protocol.PlayState, 0x5E, clientPlay.SoundEffect{})
	registry.Register(protocol.PlayState, 0x5F, clientPlay.StopSound{})
	registry.Register(protocol.PlayState, 0x60, clientPlay.SystemChatMessage{})
	registry.Register(protocol.PlayState, 0x61, clientPlay.SetTabListHeaderAndFooter{})
	registry.Register(protocol.PlayState, 0x62, clientPlay.TagQueryResponse{})
	registry.Register(protocol.PlayState, 0x63, clientPlay.PickupItem{})
	registry.Register(protocol.PlayState, 0x64, clientPlay.TeleportEntity{})
	registry.Register(protocol.PlayState, 0x65, clientPlay.UpdateAttributes{})
	registry.Register(protocol.PlayState, 0x66, clientPlay.UpdateAttributes{})
	registry.Register(protocol.PlayState, 0x67, clientPlay.FeatureFlags{})
	registry.Register(protocol.PlayState, 0x68, clientPlay.EntityEffect{})
	registry.Register(protocol.PlayState, 0x69, clientPlay.UpdateRecipes{})
	registry.Register(protocol.PlayState, 0x6A, clientPlay.UpdateTags{})

	return registry
}

func serverBound() protocol.PacketRegistry {

	registry := protocol.NewPacketRegistry()

	// Handshaking
	registry.Register(protocol.HandshakingState, 0x00, serverHandshake.Handshake{})
	registry.Register(protocol.HandshakingState, 0xFE, serverHandshake.LegacyServerListPing{})

	// Status
	registry.Register(protocol.StatusState, 0x00, serverStatus.Request{})
	registry.Register(protocol.StatusState, 0x01, serverStatus.PingRequest{})

	// Login
	registry.Register(protocol.LoginState, 0x00, serverLogin.Start{})
	registry.Register(protocol.LoginState, 0x01, serverLogin.EncryptionResponse{})
	registry.Register(protocol.LoginState, 0x02, serverLogin.PluginResponse{})

	// Play
	registry.Register(protocol.PlayState, 0x00, serverPlay.ConfirmTeleportation{})
	registry.Register(protocol.PlayState, 0x01, serverPlay.QueryBlockEntityTag{})
	registry.Register(protocol.PlayState, 0x02, serverPlay.ChangeDifficulty{})
	registry.Register(protocol.PlayState, 0x03, serverPlay.MessageAcknowledgement{})
	registry.Register(protocol.PlayState, 0x04, serverPlay.ChatCommand{})
	registry.Register(protocol.PlayState, 0x05, serverPlay.ChatMessage{})
	registry.Register(protocol.PlayState, 0x06, serverPlay.ClientCommand{})
	registry.Register(protocol.PlayState, 0x07, serverPlay.ClientInformation{})
	registry.Register(protocol.PlayState, 0x08, serverPlay.CommandSuggestionsRequest{})
	registry.Register(protocol.PlayState, 0x09, serverPlay.ClickContainerButton{})
	registry.Register(protocol.PlayState, 0x0A, serverPlay.ClickContainer{})
	registry.Register(protocol.PlayState, 0x0B, serverPlay.CloseContainer{})
	registry.Register(protocol.PlayState, 0x0C, serverPlay.PluginMessage{})
	registry.Register(protocol.PlayState, 0x0D, serverPlay.EditBook{})
	registry.Register(protocol.PlayState, 0x0E, serverPlay.QueryEntityTag{})
	registry.Register(protocol.PlayState, 0x0F, serverPlay.InteractEntity{})
	registry.Register(protocol.PlayState, 0x10, serverPlay.JigsawGenerate{})
	registry.Register(protocol.PlayState, 0x11, serverPlay.KeepAlive{})
	registry.Register(protocol.PlayState, 0x12, serverPlay.LockDifficulty{})
	registry.Register(protocol.PlayState, 0x13, serverPlay.SetPlayerPosition{})
	registry.Register(protocol.PlayState, 0x14, serverPlay.SetPlayerPositionAndRotation{})
	registry.Register(protocol.PlayState, 0x15, serverPlay.SetPlayerRotation{})
	registry.Register(protocol.PlayState, 0x16, serverPlay.SetPlayerOnGround{})
	registry.Register(protocol.PlayState, 0x17, serverPlay.MoveVehicle{})
	registry.Register(protocol.PlayState, 0x18, serverPlay.PaddleBoat{})
	registry.Register(protocol.PlayState, 0x19, serverPlay.PickItem{})
	registry.Register(protocol.PlayState, 0x1A, serverPlay.PlaceRecipe{})
	registry.Register(protocol.PlayState, 0x1B, serverPlay.PlayerAbilities{})
	registry.Register(protocol.PlayState, 0x1C, serverPlay.PlayerAction{})
	registry.Register(protocol.PlayState, 0x1D, serverPlay.PlayerCommand{})
	registry.Register(protocol.PlayState, 0x1E, serverPlay.PlayerInput{})
	registry.Register(protocol.PlayState, 0x1F, serverPlay.Pong{})
	registry.Register(protocol.PlayState, 0x20, serverPlay.PlayerSession{})
	registry.Register(protocol.PlayState, 0x21, serverPlay.ChangeRecipeBookSettings{})
	registry.Register(protocol.PlayState, 0x22, serverPlay.SetSeenRecipe{})
	registry.Register(protocol.PlayState, 0x23, serverPlay.RenameItem{})
	registry.Register(protocol.PlayState, 0x24, serverPlay.ResourcePack{})
	registry.Register(protocol.PlayState, 0x25, serverPlay.SeenAdvancements{})
	registry.Register(protocol.PlayState, 0x26, serverPlay.SelectTrade{})
	registry.Register(protocol.PlayState, 0x27, serverPlay.SetBeaconEffect{})
	registry.Register(protocol.PlayState, 0x28, serverPlay.SetHeldItem{})
	registry.Register(protocol.PlayState, 0x29, serverPlay.ProgramCommandBlock{})
	registry.Register(protocol.PlayState, 0x2A, serverPlay.ProgramCommandBlockMinecart{})
	registry.Register(protocol.PlayState, 0x2B, serverPlay.SetCreativeModeSlot{})
	registry.Register(protocol.PlayState, 0x2C, serverPlay.ProgramJigsawBlock{})
	registry.Register(protocol.PlayState, 0x2D, serverPlay.ProgramStructureBlock{})
	registry.Register(protocol.PlayState, 0x2E, serverPlay.UpdateSign{})
	registry.Register(protocol.PlayState, 0x2F, serverPlay.SwingArm{})
	registry.Register(protocol.PlayState, 0x30, serverPlay.TeleportToEntity{})
	registry.Register(protocol.PlayState, 0x31, serverPlay.UseItemOn{})
	registry.Register(protocol.PlayState, 0x32, serverPlay.UseItem{})

	return registry
}
