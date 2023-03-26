package version119

import (
	clientLogin "javaEdition/clientbound/login"
	clientPlay "javaEdition/clientbound/play"
	clientStatus "javaEdition/clientbound/status"
	"javaEdition/protocol/common"
	serverHandshake "javaEdition/serverbound/handshake"
	serverLogin "javaEdition/serverbound/login"
	serverPlay "javaEdition/serverbound/play"
	serverStatus "javaEdition/serverbound/status"
)

func Registry(bound common.Bound) common.PacketRegistry {

	switch bound {
	case common.ClientBound:
		return clientBound()
	case common.ServerBound:
		return serverBound()
	}

	return nil
}

func clientBound() common.PacketRegistry {

	registry := common.NewPacketRegistry()

	// Status
	registry.Register(common.StatusState, 0x00, clientStatus.Response{})
	registry.Register(common.StatusState, 0x01, clientStatus.PingResponse{})

	// Login
	registry.Register(common.LoginState, 0x00, clientLogin.Disconnect{})
	registry.Register(common.LoginState, 0x01, clientLogin.EncryptionRequest{})
	registry.Register(common.LoginState, 0x02, clientLogin.Success{})
	registry.Register(common.LoginState, 0x03, clientLogin.SetCompression{})
	registry.Register(common.LoginState, 0x04, clientLogin.PluginRequest{})

	// Play
	registry.Register(common.PlayState, 0x00, clientPlay.SpawnEntity{})
	registry.Register(common.PlayState, 0x01, clientPlay.SpawnExperienceOrb{})
	registry.Register(common.PlayState, 0x02, clientPlay.SpawnPlayer{})
	registry.Register(common.PlayState, 0x03, clientPlay.EntityAnimation{})
	registry.Register(common.PlayState, 0x04, clientPlay.AwardStatistics{})
	registry.Register(common.PlayState, 0x05, clientPlay.AcknowledgeBlockChange{})
	registry.Register(common.PlayState, 0x06, clientPlay.SetBlockDestroyStage{})
	registry.Register(common.PlayState, 0x07, clientPlay.BlockEntityData{})
	registry.Register(common.PlayState, 0x08, clientPlay.BlockAction{})
	registry.Register(common.PlayState, 0x09, clientPlay.BlockUpdate{})
	registry.Register(common.PlayState, 0x0A, clientPlay.BossBar{})
	registry.Register(common.PlayState, 0x0B, clientPlay.ChangeDifficulty{})
	registry.Register(common.PlayState, 0x0C, clientPlay.ClearTitles{})
	registry.Register(common.PlayState, 0x0D, clientPlay.CommandSuggestionsResponse{})
	registry.Register(common.PlayState, 0x0E, clientPlay.Commands{})
	registry.Register(common.PlayState, 0x0F, clientPlay.CloseContainer{})
	registry.Register(common.PlayState, 0x10, clientPlay.SetContainerContent{})
	registry.Register(common.PlayState, 0x11, clientPlay.SetContainerProperty{})
	registry.Register(common.PlayState, 0x12, clientPlay.SetContainerSlot{})
	registry.Register(common.PlayState, 0x13, clientPlay.SetCooldown{})
	registry.Register(common.PlayState, 0x14, clientPlay.ChatSuggestions{})
	registry.Register(common.PlayState, 0x15, clientPlay.PluginMessage{})
	registry.Register(common.PlayState, 0x16, clientPlay.DeleteMessage{})
	registry.Register(common.PlayState, 0x17, clientPlay.Disconnect{})
	registry.Register(common.PlayState, 0x18, clientPlay.DisguisedChatMessage{})
	registry.Register(common.PlayState, 0x19, clientPlay.EntityEvent{})
	registry.Register(common.PlayState, 0x1A, clientPlay.Explosion{})
	registry.Register(common.PlayState, 0x1B, clientPlay.UnloadChunk{})
	registry.Register(common.PlayState, 0x1C, clientPlay.GameEvent{})
	registry.Register(common.PlayState, 0x1D, clientPlay.OpenHorseScreen{})
	registry.Register(common.PlayState, 0x1E, clientPlay.InitializeWorldBorder{})
	registry.Register(common.PlayState, 0x1F, clientPlay.KeepAlive{})
	registry.Register(common.PlayState, 0x20, clientPlay.ChunkDataAndUpdateLight{})
	registry.Register(common.PlayState, 0x21, clientPlay.WorldEvent{})
	registry.Register(common.PlayState, 0x22, clientPlay.Particle{})
	registry.Register(common.PlayState, 0x23, clientPlay.UpdateLight{})
	registry.Register(common.PlayState, 0x24, clientPlay.Login{})
	registry.Register(common.PlayState, 0x25, clientPlay.MapData{})
	registry.Register(common.PlayState, 0x26, clientPlay.MerchantOffers{})
	registry.Register(common.PlayState, 0x27, clientPlay.UpdateEntityPosition{})
	registry.Register(common.PlayState, 0x28, clientPlay.UpdateEntityPositionAndRotation{})
	registry.Register(common.PlayState, 0x29, clientPlay.UpdateEntityRotation{})
	registry.Register(common.PlayState, 0x2A, clientPlay.MoveVehicle{})
	registry.Register(common.PlayState, 0x2B, clientPlay.OpenBook{})
	registry.Register(common.PlayState, 0x2C, clientPlay.OpenWindow{})
	registry.Register(common.PlayState, 0x2D, clientPlay.OpenSignEditor{})
	registry.Register(common.PlayState, 0x2E, clientPlay.Ping{})
	registry.Register(common.PlayState, 0x2F, clientPlay.PlaceGhostRecipe{})
	registry.Register(common.PlayState, 0x30, clientPlay.PlayerAbilities{})
	registry.Register(common.PlayState, 0x31, clientPlay.PlayerChatMessage{})
	registry.Register(common.PlayState, 0x32, clientPlay.EndCombat{})
	registry.Register(common.PlayState, 0x33, clientPlay.EnterCombat{})
	registry.Register(common.PlayState, 0x34, clientPlay.CombatDeath{})
	registry.Register(common.PlayState, 0x35, clientPlay.PlayerInfoRemove{})
	registry.Register(common.PlayState, 0x36, clientPlay.PlayerInfoUpdate{})
	registry.Register(common.PlayState, 0x37, clientPlay.LookAt{})
	registry.Register(common.PlayState, 0x38, clientPlay.SynchronizePlayerPosition{})
	registry.Register(common.PlayState, 0x39, clientPlay.UpdateRecipeBook{})
	registry.Register(common.PlayState, 0x3A, clientPlay.RemoveEntities{})
	registry.Register(common.PlayState, 0x3B, clientPlay.RemoveEntityEffect{})
	registry.Register(common.PlayState, 0x3C, clientPlay.ResourcePack{})
	registry.Register(common.PlayState, 0x3D, clientPlay.Respawn{})
	registry.Register(common.PlayState, 0x3E, clientPlay.SetHeadRotation{})
	registry.Register(common.PlayState, 0x3F, clientPlay.UpdateSectionBlocks{})
	registry.Register(common.PlayState, 0x40, clientPlay.SelectAdvancementTab{})
	registry.Register(common.PlayState, 0x41, clientPlay.ServerData{})
	registry.Register(common.PlayState, 0x42, clientPlay.SetActionBarText{})
	registry.Register(common.PlayState, 0x43, clientPlay.SetBorderCenter{})
	registry.Register(common.PlayState, 0x44, clientPlay.SetBorderLerpSize{})
	registry.Register(common.PlayState, 0x45, clientPlay.SetBorderSize{})
	registry.Register(common.PlayState, 0x46, clientPlay.SetBorderWarningDelay{})
	registry.Register(common.PlayState, 0x47, clientPlay.SetBorderWarningDistance{})
	registry.Register(common.PlayState, 0x48, clientPlay.SetCamera{})
	registry.Register(common.PlayState, 0x49, clientPlay.SetHeldItem{})
	registry.Register(common.PlayState, 0x4A, clientPlay.SetCenterChunk{})
	registry.Register(common.PlayState, 0x4B, clientPlay.SetRenderDistance{})
	registry.Register(common.PlayState, 0x4C, clientPlay.SetDefaultSpawnPosition{})
	registry.Register(common.PlayState, 0x4D, clientPlay.DisplayObjective{})
	registry.Register(common.PlayState, 0x4E, clientPlay.SetEntityMetadata{})
	registry.Register(common.PlayState, 0x4F, clientPlay.LinkEntities{})
	registry.Register(common.PlayState, 0x50, clientPlay.SetEntityVelocity{})
	registry.Register(common.PlayState, 0x51, clientPlay.SetEquipment{})
	registry.Register(common.PlayState, 0x52, clientPlay.SetExperience{})
	registry.Register(common.PlayState, 0x53, clientPlay.SetHealth{})
	registry.Register(common.PlayState, 0x54, clientPlay.UpdateObjectives{})
	registry.Register(common.PlayState, 0x55, clientPlay.SetPassengers{})
	registry.Register(common.PlayState, 0x56, clientPlay.UpdateTeams{})
	registry.Register(common.PlayState, 0x57, clientPlay.UpdateScore{})
	registry.Register(common.PlayState, 0x58, clientPlay.SetSimulationDistance{})
	registry.Register(common.PlayState, 0x59, clientPlay.SetSubtitleText{})
	registry.Register(common.PlayState, 0x5A, clientPlay.UpdateTime{})
	registry.Register(common.PlayState, 0x5B, clientPlay.SetTitleText{})
	registry.Register(common.PlayState, 0x5C, clientPlay.SetTitleAnimationTimes{})
	registry.Register(common.PlayState, 0x5D, clientPlay.EntitySoundEffect{})
	registry.Register(common.PlayState, 0x5E, clientPlay.SoundEffect{})
	registry.Register(common.PlayState, 0x5F, clientPlay.StopSound{})
	registry.Register(common.PlayState, 0x60, clientPlay.SystemChatMessage{})
	registry.Register(common.PlayState, 0x61, clientPlay.SetTabListHeaderAndFooter{})
	registry.Register(common.PlayState, 0x62, clientPlay.TagQueryResponse{})
	registry.Register(common.PlayState, 0x63, clientPlay.PickupItem{})
	registry.Register(common.PlayState, 0x64, clientPlay.TeleportEntity{})
	registry.Register(common.PlayState, 0x65, clientPlay.UpdateAttributes{})
	registry.Register(common.PlayState, 0x66, clientPlay.UpdateAttributes{})
	registry.Register(common.PlayState, 0x67, clientPlay.FeatureFlags{})
	registry.Register(common.PlayState, 0x68, clientPlay.EntityEffect{})
	registry.Register(common.PlayState, 0x69, clientPlay.UpdateRecipes{})
	registry.Register(common.PlayState, 0x6A, clientPlay.UpdateTags{})

	return registry
}

func serverBound() common.PacketRegistry {

	registry := common.NewPacketRegistry()

	// Handshaking
	registry.Register(common.HandshakingState, 0x00, serverHandshake.Handshake{})
	registry.Register(common.HandshakingState, 0xFE, serverHandshake.LegacyServerListPing{})

	// Status
	registry.Register(common.StatusState, 0x00, serverStatus.Request{})
	registry.Register(common.StatusState, 0x01, serverStatus.PingRequest{})

	// Login
	registry.Register(common.LoginState, 0x00, serverLogin.Start{})
	registry.Register(common.LoginState, 0x01, serverLogin.EncryptionResponse{})
	registry.Register(common.LoginState, 0x02, serverLogin.PluginResponse{})

	// Play
	registry.Register(common.PlayState, 0x00, serverPlay.ConfirmTeleportation{})
	registry.Register(common.PlayState, 0x01, serverPlay.QueryBlockEntityTag{})
	registry.Register(common.PlayState, 0x02, serverPlay.ChangeDifficulty{})
	registry.Register(common.PlayState, 0x03, serverPlay.MessageAcknowledgement{})
	registry.Register(common.PlayState, 0x04, serverPlay.ChatCommand{})
	registry.Register(common.PlayState, 0x05, serverPlay.ChatMessage{})
	registry.Register(common.PlayState, 0x06, serverPlay.ClientCommand{})
	registry.Register(common.PlayState, 0x07, serverPlay.ClientInformation{})
	registry.Register(common.PlayState, 0x08, serverPlay.CommandSuggestionsRequest{})
	registry.Register(common.PlayState, 0x09, serverPlay.ClickContainerButton{})
	registry.Register(common.PlayState, 0x0A, serverPlay.ClickContainer{})
	registry.Register(common.PlayState, 0x0B, serverPlay.CloseContainer{})
	registry.Register(common.PlayState, 0x0C, serverPlay.PluginMessage{})
	registry.Register(common.PlayState, 0x0D, serverPlay.EditBook{})
	registry.Register(common.PlayState, 0x0E, serverPlay.QueryEntityTag{})
	registry.Register(common.PlayState, 0x0F, serverPlay.InteractEntity{})
	registry.Register(common.PlayState, 0x10, serverPlay.JigsawGenerate{})
	registry.Register(common.PlayState, 0x11, serverPlay.KeepAlive{})
	registry.Register(common.PlayState, 0x12, serverPlay.LockDifficulty{})
	registry.Register(common.PlayState, 0x13, serverPlay.SetPlayerPosition{})
	registry.Register(common.PlayState, 0x14, serverPlay.SetPlayerPositionAndRotation{})
	registry.Register(common.PlayState, 0x15, serverPlay.SetPlayerRotation{})
	registry.Register(common.PlayState, 0x16, serverPlay.SetPlayerOnGround{})
	registry.Register(common.PlayState, 0x17, serverPlay.MoveVehicle{})
	registry.Register(common.PlayState, 0x18, serverPlay.PaddleBoat{})
	registry.Register(common.PlayState, 0x19, serverPlay.PickItem{})
	registry.Register(common.PlayState, 0x1A, serverPlay.PlaceRecipe{})
	registry.Register(common.PlayState, 0x1B, serverPlay.PlayerAbilities{})
	registry.Register(common.PlayState, 0x1C, serverPlay.PlayerAction{})
	registry.Register(common.PlayState, 0x1D, serverPlay.PlayerCommand{})
	registry.Register(common.PlayState, 0x1E, serverPlay.PlayerInput{})
	registry.Register(common.PlayState, 0x1F, serverPlay.Pong{})
	registry.Register(common.PlayState, 0x20, serverPlay.PlayerSession{})
	registry.Register(common.PlayState, 0x21, serverPlay.ChangeRecipeBookSettings{})
	registry.Register(common.PlayState, 0x22, serverPlay.SetSeenRecipe{})
	registry.Register(common.PlayState, 0x23, serverPlay.RenameItem{})
	registry.Register(common.PlayState, 0x24, serverPlay.ResourcePack{})
	registry.Register(common.PlayState, 0x25, serverPlay.SeenAdvancements{})
	registry.Register(common.PlayState, 0x26, serverPlay.SelectTrade{})
	registry.Register(common.PlayState, 0x27, serverPlay.SetBeaconEffect{})
	registry.Register(common.PlayState, 0x28, serverPlay.SetHeldItem{})
	registry.Register(common.PlayState, 0x29, serverPlay.ProgramCommandBlock{})
	registry.Register(common.PlayState, 0x2A, serverPlay.ProgramCommandBlockMinecart{})
	registry.Register(common.PlayState, 0x2B, serverPlay.SetCreativeModeSlot{})
	registry.Register(common.PlayState, 0x2C, serverPlay.ProgramJigsawBlock{})
	registry.Register(common.PlayState, 0x2D, serverPlay.ProgramStructureBlock{})
	registry.Register(common.PlayState, 0x2E, serverPlay.UpdateSign{})
	registry.Register(common.PlayState, 0x2F, serverPlay.SwingArm{})
	registry.Register(common.PlayState, 0x30, serverPlay.TeleportToEntity{})
	registry.Register(common.PlayState, 0x31, serverPlay.UseItemOn{})
	registry.Register(common.PlayState, 0x32, serverPlay.UseItem{})

	return registry
}
