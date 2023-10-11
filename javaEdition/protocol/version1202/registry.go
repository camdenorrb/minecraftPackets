package version1202

import (
	"github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"
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
	registry.Register(common.StatusState, 0x00, string(common.ClientBoundStatusResponse))
	registry.Register(common.StatusState, 0x01, string(common.ClientBoundStatusPong))

	// Configuration
	registry.Register(common.ConfigurationState, 0x00, string(common.ClientBoundConfigurationPluginMessage))
	registry.Register(common.ConfigurationState, 0x01, string(common.ClientBoundConfigurationDisconnect))
	registry.Register(common.ConfigurationState, 0x02, string(common.ClientBoundConfigurationFinish))
	registry.Register(common.ConfigurationState, 0x03, string(common.ClientBoundConfigurationKeepAlive))
	registry.Register(common.ConfigurationState, 0x04, string(common.ClientBoundConfigurationPing))
	registry.Register(common.ConfigurationState, 0x05, string(common.ClientBoundConfigurationRegistryData))
	registry.Register(common.ConfigurationState, 0x06, string(common.ClientBoundConfigurationResourcePack))
	registry.Register(common.ConfigurationState, 0x07, string(common.ClientBoundConfigurationFeatureFlags))
	registry.Register(common.ConfigurationState, 0x08, string(common.ClientBoundConfigurationUpdateTags))

	// Login
	registry.Register(common.LoginState, 0x00, string(common.ClientBoundLoginDisconnect))
	registry.Register(common.LoginState, 0x01, string(common.ClientBoundLoginEncryptionRequest))
	registry.Register(common.LoginState, 0x02, string(common.ClientBoundLoginSuccess))
	registry.Register(common.LoginState, 0x03, string(common.ClientBoundLoginSetCompression))
	registry.Register(common.LoginState, 0x04, string(common.ClientBoundLoginPluginRequest))

	// Play
	registry.Register(common.PlayState, 0x00, string(common.ClientBoundPlayBundleDelimiter))
	registry.Register(common.PlayState, 0x01, string(common.ClientBoundPlaySpawnEntity))
	registry.Register(common.PlayState, 0x02, string(common.ClientBoundPlaySpawnExperienceOrb))
	registry.Register(common.PlayState, 0x03, string(common.ClientBoundPlayEntityAnimation))
	registry.Register(common.PlayState, 0x04, string(common.ClientBoundPlayAwardStatistics))
	registry.Register(common.PlayState, 0x05, string(common.ClientBoundPlayAcknowledgeBlockChange))
	registry.Register(common.PlayState, 0x06, string(common.ClientBoundPlaySetBlockDestroyStage))
	registry.Register(common.PlayState, 0x07, string(common.ClientBoundPlayBlockEntityData))
	registry.Register(common.PlayState, 0x08, string(common.ClientBoundPlayBlockAction))
	registry.Register(common.PlayState, 0x09, string(common.ClientBoundPlayBlockUpdate))
	registry.Register(common.PlayState, 0x0A, string(common.ClientBoundPlayBossBar))
	registry.Register(common.PlayState, 0x0B, string(common.ClientBoundPlayChangeDifficulty))
	registry.Register(common.PlayState, 0x0C, string(common.ClientBoundPlayChunkBatchFinished))
	registry.Register(common.PlayState, 0x0D, string(common.ClientBoundPlayChunkBatchStart))
	registry.Register(common.PlayState, 0x0E, string(common.ClientBoundPlayChunkBiomes))
	registry.Register(common.PlayState, 0x0F, string(common.ClientBoundPlayClearTitles))
	registry.Register(common.PlayState, 0x10, string(common.ClientBoundPlayCommandSuggestionsResponse))
	registry.Register(common.PlayState, 0x11, string(common.ClientBoundPlayCommands))
	registry.Register(common.PlayState, 0x12, string(common.ClientBoundPlayCloseContainer))
	registry.Register(common.PlayState, 0x13, string(common.ClientBoundPlaySetContainerContent))
	registry.Register(common.PlayState, 0x14, string(common.ClientBoundPlaySetContainerProperty))
	registry.Register(common.PlayState, 0x15, string(common.ClientBoundPlaySetContainerSlot))
	registry.Register(common.PlayState, 0x16, string(common.ClientBoundPlaySetCooldown))
	registry.Register(common.PlayState, 0x17, string(common.ClientBoundPlayChatSuggestions))
	registry.Register(common.PlayState, 0x18, string(common.ClientBoundPlayPluginMessage))
	registry.Register(common.PlayState, 0x19, string(common.ClientBoundPlayDamageEvent))
	registry.Register(common.PlayState, 0x1A, string(common.ClientBoundPlayDeleteMessage))
	registry.Register(common.PlayState, 0x1B, string(common.ClientBoundPlayDisconnect))
	registry.Register(common.PlayState, 0x1C, string(common.ClientBoundPlayDisguisedChatMessage))
	registry.Register(common.PlayState, 0x1D, string(common.ClientBoundPlayEntityEvent))
	registry.Register(common.PlayState, 0x1E, string(common.ClientBoundPlayExplosion))
	registry.Register(common.PlayState, 0x1F, string(common.ClientBoundPlayUnloadChunk))
	registry.Register(common.PlayState, 0x20, string(common.ClientBoundPlayGameEvent))
	registry.Register(common.PlayState, 0x21, string(common.ClientBoundPlayOpenHorseScreen))
	registry.Register(common.PlayState, 0x22, string(common.ClientBoundPlayHurtAnimation))
	registry.Register(common.PlayState, 0x23, string(common.ClientBoundPlayInitializeWorldBorder))
	registry.Register(common.PlayState, 0x24, string(common.ClientBoundPlayKeepAlive))
	registry.Register(common.PlayState, 0x25, string(common.ClientBoundPlayChunkDataAndUpdateLight))
	registry.Register(common.PlayState, 0x26, string(common.ClientBoundPlayWorldEvent))
	registry.Register(common.PlayState, 0x27, string(common.ClientBoundPlayParticle))
	registry.Register(common.PlayState, 0x28, string(common.ClientBoundPlayUpdateLight))
	registry.Register(common.PlayState, 0x29, string(common.ClientBoundPlayLogin))
	registry.Register(common.PlayState, 0x2A, string(common.ClientBoundPlayMapData))
	registry.Register(common.PlayState, 0x2B, string(common.ClientBoundPlayMerchantOffers))
	registry.Register(common.PlayState, 0x2C, string(common.ClientBoundPlayUpdateEntityPosition))
	registry.Register(common.PlayState, 0x2D, string(common.ClientBoundPlayUpdateEntityPositionAndRotation))
	registry.Register(common.PlayState, 0x2E, string(common.ClientBoundPlayUpdateEntityRotation))
	registry.Register(common.PlayState, 0x2F, string(common.ClientBoundPlayMoveVehicle))
	registry.Register(common.PlayState, 0x30, string(common.ClientBoundPlayOpenBook))
	registry.Register(common.PlayState, 0x31, string(common.ClientBoundPlayOpenScreen))
	registry.Register(common.PlayState, 0x32, string(common.ClientBoundPlayOpenSignEditor))
	registry.Register(common.PlayState, 0x33, string(common.ClientBoundPlayPing))
	registry.Register(common.PlayState, 0x34, string(common.ClientBoundPlayPingResponse))
	registry.Register(common.PlayState, 0x35, string(common.ClientBoundPlayPlaceGhostRecipe))
	registry.Register(common.PlayState, 0x36, string(common.ClientBoundPlayPlayerAbilities))
	registry.Register(common.PlayState, 0x37, string(common.ClientBoundPlayPlayerChatMessage))
	registry.Register(common.PlayState, 0x38, string(common.ClientBoundPlayEndCombat))
	registry.Register(common.PlayState, 0x39, string(common.ClientBoundPlayEnterCombat))
	registry.Register(common.PlayState, 0x3A, string(common.ClientBoundPlayCombatDeath))
	registry.Register(common.PlayState, 0x3B, string(common.ClientBoundPlayPlayerInfoRemove))
	registry.Register(common.PlayState, 0x3C, string(common.ClientBoundPlayPlayerInfoUpdate))
	registry.Register(common.PlayState, 0x3D, string(common.ClientBoundPlayLookAt))
	registry.Register(common.PlayState, 0x3E, string(common.ClientBoundPlaySynchronizePlayerPosition))
	registry.Register(common.PlayState, 0x3F, string(common.ClientBoundPlayUpdateRecipeBook))
	registry.Register(common.PlayState, 0x40, string(common.ClientBoundPlayRemoveEntities))
	registry.Register(common.PlayState, 0x41, string(common.ClientBoundPlayRemoveEntityEffect))
	registry.Register(common.PlayState, 0x42, string(common.ClientBoundPlayResourcePack))
	registry.Register(common.PlayState, 0x43, string(common.ClientBoundPlayRespawn))
	registry.Register(common.PlayState, 0x44, string(common.ClientBoundPlaySetHeadRotation))
	registry.Register(common.PlayState, 0x45, string(common.ClientBoundPlayUpdateSectionBlocks))
	registry.Register(common.PlayState, 0x46, string(common.ClientBoundPlaySelectAdvancementTab))
	registry.Register(common.PlayState, 0x47, string(common.ClientBoundPlayServerData))
	registry.Register(common.PlayState, 0x48, string(common.ClientBoundPlaySetActionBarText))
	registry.Register(common.PlayState, 0x49, string(common.ClientBoundPlaySetBorderCenter))
	registry.Register(common.PlayState, 0x4A, string(common.ClientBoundPlaySetBorderLerpSize))
	registry.Register(common.PlayState, 0x4B, string(common.ClientBoundPlaySetBorderSize))
	registry.Register(common.PlayState, 0x4C, string(common.ClientBoundPlaySetBorderWarningDelay))
	registry.Register(common.PlayState, 0x4D, string(common.ClientBoundPlaySetBorderWarningDistance))
	registry.Register(common.PlayState, 0x4E, string(common.ClientBoundPlaySetCamera))
	registry.Register(common.PlayState, 0x4F, string(common.ClientBoundPlaySetHeldItem))
	registry.Register(common.PlayState, 0x50, string(common.ClientBoundPlaySetCenterChunk))
	registry.Register(common.PlayState, 0x51, string(common.ClientBoundPlaySetRenderDistance))
	registry.Register(common.PlayState, 0x52, string(common.ClientBoundPlaySetDefaultSpawnPosition))
	registry.Register(common.PlayState, 0x53, string(common.ClientBoundPlayDisplayObjective))
	registry.Register(common.PlayState, 0x54, string(common.ClientBoundPlaySetEntityMetadata))
	registry.Register(common.PlayState, 0x55, string(common.ClientBoundPlayLinkEntities))
	registry.Register(common.PlayState, 0x56, string(common.ClientBoundPlaySetEntityVelocity))
	registry.Register(common.PlayState, 0x57, string(common.ClientBoundPlaySetEquipment))
	registry.Register(common.PlayState, 0x58, string(common.ClientBoundPlaySetExperience))
	registry.Register(common.PlayState, 0x59, string(common.ClientBoundPlaySetHealth))
	registry.Register(common.PlayState, 0x5A, string(common.ClientBoundPlayUpdateObjectives))
	registry.Register(common.PlayState, 0x5B, string(common.ClientBoundPlaySetPassengers))
	registry.Register(common.PlayState, 0x5C, string(common.ClientBoundPlayUpdateTeams))
	registry.Register(common.PlayState, 0x5D, string(common.ClientBoundPlayUpdateScore))
	registry.Register(common.PlayState, 0x5E, string(common.ClientBoundPlaySetSimulationDistance))
	registry.Register(common.PlayState, 0x5F, string(common.ClientBoundPlaySetSubtitleText))
	registry.Register(common.PlayState, 0x60, string(common.ClientBoundPlayUpdateTime))
	registry.Register(common.PlayState, 0x61, string(common.ClientBoundPlaySetTitleText))
	registry.Register(common.PlayState, 0x62, string(common.ClientBoundPlaySetTitleAnimationTimes))
	registry.Register(common.PlayState, 0x63, string(common.ClientBoundPlayEntitySoundEffect))
	registry.Register(common.PlayState, 0x64, string(common.ClientBoundPlaySoundEffect))
	registry.Register(common.PlayState, 0x65, string(common.ClientBoundPlayStartConfiguration))
	registry.Register(common.PlayState, 0x66, string(common.ClientBoundPlayStopSound))
	registry.Register(common.PlayState, 0x67, string(common.ClientBoundPlaySystemChatMessage))
	registry.Register(common.PlayState, 0x68, string(common.ClientBoundPlaySetTabListHeaderAndFooter))
	registry.Register(common.PlayState, 0x69, string(common.ClientBoundPlayTagQueryResponse))
	registry.Register(common.PlayState, 0x6A, string(common.ClientBoundPlayPickupItem))
	registry.Register(common.PlayState, 0x6B, string(common.ClientBoundPlayTeleportEntity))
	registry.Register(common.PlayState, 0x6C, string(common.ClientBoundPlayUpdateAdvancements))
	registry.Register(common.PlayState, 0x6D, string(common.ClientBoundPlayUpdateAttributes))
	registry.Register(common.PlayState, 0x6E, string(common.ClientBoundPlayEntityEffect))
	registry.Register(common.PlayState, 0x6F, string(common.ClientBoundPlayUpdateRecipes))
	registry.Register(common.PlayState, 0x70, string(common.ClientBoundPlayUpdateTags))

	return registry
}

func serverBound() common.PacketRegistry {

	registry := common.NewPacketRegistry()

	// Handshaking
	registry.Register(common.HandshakingState, 0x00, string(common.ServerBoundHandshake))
	registry.Register(common.HandshakingState, 0xFE, string(common.ServerBoundHandshakeLegacyServerListPing))

	// Status
	registry.Register(common.StatusState, 0x00, string(common.ServerBoundStatusRequest))
	registry.Register(common.StatusState, 0x01, string(common.ServerBoundStatusPing))

	// Configuration
	registry.Register(common.ConfigurationState, 0x00, string(common.ServerBoundConfigurationClientInformation))
	registry.Register(common.ConfigurationState, 0x01, string(common.ServerBoundConfigurationPluginMessage))
	registry.Register(common.ConfigurationState, 0x02, string(common.ServerBoundConfigurationFinish))
	registry.Register(common.ConfigurationState, 0x03, string(common.ServerBoundConfigurationKeepAlive))
	registry.Register(common.ConfigurationState, 0x04, string(common.ServerBoundConfigurationPong))
	registry.Register(common.ConfigurationState, 0x05, string(common.ServerBoundConfigurationResourcePackResponse))

	// Login
	registry.Register(common.LoginState, 0x00, string(common.ServerBoundLoginStart))
	registry.Register(common.LoginState, 0x01, string(common.ServerBoundLoginEncryptionResponse))
	registry.Register(common.LoginState, 0x02, string(common.ServerBoundLoginPluginResponse))
	registry.Register(common.LoginState, 0x03, string(common.ServerBoundLoginAcknowledged))

	// Play
	registry.Register(common.PlayState, 0x00, string(common.ServerBoundPlayConfirmTeleportation))
	registry.Register(common.PlayState, 0x01, string(common.ServerBoundPlayQueryBlockEntityTag))
	registry.Register(common.PlayState, 0x02, string(common.ServerBoundPlayChangeDifficulty))
	registry.Register(common.PlayState, 0x03, string(common.ServerBoundPlayMessageAcknowledgement))
	registry.Register(common.PlayState, 0x04, string(common.ServerBoundPlayChatCommand))
	registry.Register(common.PlayState, 0x05, string(common.ServerBoundPlayChatMessage))
	registry.Register(common.PlayState, 0x06, string(common.ServerBoundPlayPlayerSession))
	registry.Register(common.PlayState, 0x07, string(common.ServerBoundPlayChunkBatchReceived))
	registry.Register(common.PlayState, 0x08, string(common.ServerBoundPlayClientCommand))
	registry.Register(common.PlayState, 0x09, string(common.ServerBoundPlayClientInformation))
	registry.Register(common.PlayState, 0x0A, string(common.ServerBoundPlayCommandSuggestionsRequest))
	registry.Register(common.PlayState, 0x0B, string(common.ServerBoundPlayConfigurationAcknowledged))
	registry.Register(common.PlayState, 0x0C, string(common.ServerBoundPlayClickContainerButton))
	registry.Register(common.PlayState, 0x0D, string(common.ServerBoundPlayClickContainer))
	registry.Register(common.PlayState, 0x0E, string(common.ServerBoundPlayCloseContainer))
	registry.Register(common.PlayState, 0x0F, string(common.ServerBoundPlayPluginMessage))
	registry.Register(common.PlayState, 0x10, string(common.ServerBoundPlayEditBook))
	registry.Register(common.PlayState, 0x11, string(common.ServerBoundPlayQueryEntityTag))
	registry.Register(common.PlayState, 0x12, string(common.ServerBoundPlayInteractEntity))
	registry.Register(common.PlayState, 0x13, string(common.ServerBoundPlayJigsawGenerate))
	registry.Register(common.PlayState, 0x14, string(common.ServerBoundPlayKeepAlive))
	registry.Register(common.PlayState, 0x15, string(common.ServerBoundPlayLockDifficulty))
	registry.Register(common.PlayState, 0x16, string(common.ServerBoundPlaySetPlayerPosition))
	registry.Register(common.PlayState, 0x17, string(common.ServerBoundPlaySetPlayerPositionAndRotation))
	registry.Register(common.PlayState, 0x18, string(common.ServerBoundPlaySetPlayerRotation))
	registry.Register(common.PlayState, 0x19, string(common.ServerBoundPlaySetPlayerOnGround))
	registry.Register(common.PlayState, 0x1A, string(common.ServerBoundPlayMoveVehicle))
	registry.Register(common.PlayState, 0x1B, string(common.ServerBoundPlayPaddleBoat))
	registry.Register(common.PlayState, 0x1C, string(common.ServerBoundPlayPickItem))
	registry.Register(common.PlayState, 0x1D, string(common.ServerBoundPlayPingRequest))
	registry.Register(common.PlayState, 0x1E, string(common.ServerBoundPlayPlaceRecipe))
	registry.Register(common.PlayState, 0x1F, string(common.ServerBoundPlayPlayerAbilities))
	registry.Register(common.PlayState, 0x20, string(common.ServerBoundPlayPlayerAction))
	registry.Register(common.PlayState, 0x21, string(common.ServerBoundPlayPlayerCommand))
	registry.Register(common.PlayState, 0x22, string(common.ServerBoundPlayPlayerInput))
	registry.Register(common.PlayState, 0x23, string(common.ServerBoundPlayPong))
	registry.Register(common.PlayState, 0x24, string(common.ServerBoundPlayChangeRecipeBookSettings))
	registry.Register(common.PlayState, 0x25, string(common.ServerBoundPlaySetSeenRecipe))
	registry.Register(common.PlayState, 0x26, string(common.ServerBoundPlayRenameItem))
	registry.Register(common.PlayState, 0x27, string(common.ServerBoundPlayResourcePack))
	registry.Register(common.PlayState, 0x28, string(common.ServerBoundPlaySeenAdvancements))
	registry.Register(common.PlayState, 0x29, string(common.ServerBoundPlaySelectTrade))
	registry.Register(common.PlayState, 0x2A, string(common.ServerBoundPlaySetBeaconEffect))
	registry.Register(common.PlayState, 0x2B, string(common.ServerBoundPlaySetHeldItem))
	registry.Register(common.PlayState, 0x2C, string(common.ServerBoundPlayProgramCommandBlock))
	registry.Register(common.PlayState, 0x2D, string(common.ServerBoundPlayProgramCommandBlockMinecart))
	registry.Register(common.PlayState, 0x2E, string(common.ServerBoundPlaySetCreativeModeSlot))
	registry.Register(common.PlayState, 0x2F, string(common.ServerBoundPlayProgramJigsawBlock))
	registry.Register(common.PlayState, 0x30, string(common.ServerBoundPlayProgramStructureBlock))
	registry.Register(common.PlayState, 0x31, string(common.ServerBoundPlayUpdateSign))
	registry.Register(common.PlayState, 0x32, string(common.ServerBoundPlaySwingArm))
	registry.Register(common.PlayState, 0x33, string(common.ServerBoundPlayTeleportToEntity))
	registry.Register(common.PlayState, 0x34, string(common.ServerBoundPlayUseItemOn))
	registry.Register(common.PlayState, 0x35, string(common.ServerBoundPlayUseItem))

	return registry
}
