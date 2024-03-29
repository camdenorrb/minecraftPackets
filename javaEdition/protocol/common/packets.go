package common

type ClientBoundPacket string

const (

	// Status
	ClientBoundStatusResponse ClientBoundPacket = "StatusResponse"
	ClientBoundStatusPong     ClientBoundPacket = "StatusPong"

	// Configuration
	ClientBoundConfigurationBrand           ClientBoundPacket = "PluginMessage"
	ClientBoundConfigurationDisconnect      ClientBoundPacket = "Disconnect"
	ClientBoundConfigurationFinish          ClientBoundPacket = "Finish"
	ClientBoundConfigurationKeepAlive       ClientBoundPacket = "KeepAlive"
	ClientBoundConfigurationPing            ClientBoundPacket = "Ping"
	ClientBoundConfigurationRegistryData    ClientBoundPacket = "RegistryData"
	ClientBoundConfigurationRemResourcePack ClientBoundPacket = "RemResourcePack"
	ClientBoundConfigurationAddResourcePack ClientBoundPacket = "AddResourcePack"
	ClientBoundConfigurationFeatureFlags    ClientBoundPacket = "FeatureFlags"
	ClientBoundConfigurationUpdateTags      ClientBoundPacket = "UpdateTags"

	// Login
	ClientBoundLoginDisconnect        ClientBoundPacket = "Disconnect"
	ClientBoundLoginEncryptionRequest ClientBoundPacket = "EncryptionRequest"
	ClientBoundLoginSuccess           ClientBoundPacket = "LoginSuccess"
	ClientBoundLoginSetCompression    ClientBoundPacket = "SetCompression"
	ClientBoundLoginPluginRequest     ClientBoundPacket = "LoginPluginRequest"

	// Play
	ClientBoundPlayBundleDelimiter                 ClientBoundPacket = "BundleDelimiter"
	ClientBoundPlaySpawnEntity                     ClientBoundPacket = "SpawnEntity"
	ClientBoundPlaySpawnExperienceOrb              ClientBoundPacket = "SpawnExperienceOrb"
	ClientBoundPlaySpawnPlayer                     ClientBoundPacket = "SpawnPlayer"
	ClientBoundPlayEntityAnimation                 ClientBoundPacket = "EntityAnimation"
	ClientBoundPlayAwardStatistics                 ClientBoundPacket = "AwardStatistics"
	ClientBoundPlayAcknowledgeBlockChange          ClientBoundPacket = "AcknowledgeBlockChange"
	ClientBoundPlaySetBlockDestroyStage            ClientBoundPacket = "SetBlockDestroyStage"
	ClientBoundPlayBlockEntityData                 ClientBoundPacket = "BlockEntityData"
	ClientBoundPlayBlockAction                     ClientBoundPacket = "BlockAction"
	ClientBoundPlayBlockUpdate                     ClientBoundPacket = "BlockUpdate"
	ClientBoundPlayBossBar                         ClientBoundPacket = "BossBar"
	ClientBoundPlayChangeDifficulty                ClientBoundPacket = "ChangeDifficulty"
	ClientBoundPlayChunkBatchFinished              ClientBoundPacket = "ChunkBatchFinished"
	ClientBoundPlayChunkBatchStart                 ClientBoundPacket = "ChunkBatchStart"
	ClientBoundPlayChunkBiomes                     ClientBoundPacket = "ChunkBiomes"
	ClientBoundPlayClearTitles                     ClientBoundPacket = "ClearTitles"
	ClientBoundPlayCommandSuggestionsResponse      ClientBoundPacket = "CommandSuggestionsResponse"
	ClientBoundPlayCommands                        ClientBoundPacket = "Commands"
	ClientBoundPlayCloseContainer                  ClientBoundPacket = "CloseContainer"
	ClientBoundPlaySetContainerContent             ClientBoundPacket = "SetContainerContent"
	ClientBoundPlaySetContainerProperty            ClientBoundPacket = "SetContainerProperty"
	ClientBoundPlaySetContainerSlot                ClientBoundPacket = "SetContainerSlot"
	ClientBoundPlaySetCooldown                     ClientBoundPacket = "SetCooldown"
	ClientBoundPlayChatSuggestions                 ClientBoundPacket = "ChatSuggestions"
	ClientBoundPlayPluginMessage                   ClientBoundPacket = "PluginMessage"
	ClientBoundPlayDamageEvent                     ClientBoundPacket = "DamageEvent"
	ClientBoundPlayDeleteMessage                   ClientBoundPacket = "DeleteMessage"
	ClientBoundPlayDisconnect                      ClientBoundPacket = "Disconnect"
	ClientBoundPlayDisguisedChatMessage            ClientBoundPacket = "DisguisedChatMessage"
	ClientBoundPlayEntityEvent                     ClientBoundPacket = "EntityEvent"
	ClientBoundPlayExplosion                       ClientBoundPacket = "Explosion"
	ClientBoundPlayUnloadChunk                     ClientBoundPacket = "UnloadChunk"
	ClientBoundPlayGameEvent                       ClientBoundPacket = "GameEvent"
	ClientBoundPlayOpenHorseScreen                 ClientBoundPacket = "OpenHorseScreen"
	ClientBoundPlayHurtAnimation                   ClientBoundPacket = "HurtAnimation"
	ClientBoundPlayInitializeWorldBorder           ClientBoundPacket = "InitializeWorldBorder"
	ClientBoundPlayKeepAlive                       ClientBoundPacket = "KeepAlive"
	ClientBoundPlayChunkDataAndUpdateLight         ClientBoundPacket = "ChunkDataAndUpdateLight"
	ClientBoundPlayWorldEvent                      ClientBoundPacket = "WorldEvent"
	ClientBoundPlayParticle                        ClientBoundPacket = "Particle"
	ClientBoundPlayUpdateLight                     ClientBoundPacket = "UpdateLight"
	ClientBoundPlayLogin                           ClientBoundPacket = "Login"
	ClientBoundPlayMapData                         ClientBoundPacket = "MapData"
	ClientBoundPlayMerchantOffers                  ClientBoundPacket = "MerchantOffers"
	ClientBoundPlayUpdateEntityPosition            ClientBoundPacket = "UpdateEntityPosition"
	ClientBoundPlayUpdateEntityPositionAndRotation ClientBoundPacket = "UpdateEntityPositionAndRotation"
	ClientBoundPlayUpdateEntityRotation            ClientBoundPacket = "UpdateEntityRotation"
	ClientBoundPlayMoveVehicle                     ClientBoundPacket = "MoveVehicle"
	ClientBoundPlayOpenBook                        ClientBoundPacket = "OpenBook"
	ClientBoundPlayOpenScreen                      ClientBoundPacket = "OpenScreen"
	ClientBoundPlayOpenSignEditor                  ClientBoundPacket = "OpenSignEditor"
	ClientBoundPlayPing                            ClientBoundPacket = "Ping"
	ClientBoundPlayPingResponse                    ClientBoundPacket = "PingResponse"
	ClientBoundPlayPlaceGhostRecipe                ClientBoundPacket = "PlaceGhostRecipe"
	ClientBoundPlayPlayerAbilities                 ClientBoundPacket = "PlayerAbilities"
	ClientBoundPlayPlayerChatMessage               ClientBoundPacket = "PlayerChatMessage"
	ClientBoundPlayEndCombat                       ClientBoundPacket = "EndCombat"
	ClientBoundPlayEnterCombat                     ClientBoundPacket = "EnterCombat"
	ClientBoundPlayCombatDeath                     ClientBoundPacket = "CombatDeath"
	ClientBoundPlayPlayerInfoRemove                ClientBoundPacket = "PlayerInfoRemove"
	ClientBoundPlayPlayerInfoUpdate                ClientBoundPacket = "PlayerInfoUpdate"
	ClientBoundPlayLookAt                          ClientBoundPacket = "LookAt"
	ClientBoundPlaySynchronizePlayerPosition       ClientBoundPacket = "SynchronizePlayerPosition"
	ClientBoundPlayUpdateRecipeBook                ClientBoundPacket = "UpdateRecipeBook"
	ClientBoundPlayRemoveEntities                  ClientBoundPacket = "RemoveEntities"
	ClientBoundPlayRemoveEntityEffect              ClientBoundPacket = "RemoveEntityEffect"
	ClientBoundPlayResourcePack                    ClientBoundPacket = "ResourcePack"
	ClientBoundPlayRespawn                         ClientBoundPacket = "Respawn"
	ClientBoundPlayResetScore                      ClientBoundPacket = "ResetScore"
	ClientBoundPlayRemoveResourcePack              ClientBoundPacket = "RemoveResourcePack"
	ClientBoundPlaySetHeadRotation                 ClientBoundPacket = "SetHeadRotation"
	ClientBoundPlayUpdateSectionBlocks             ClientBoundPacket = "UpdateSectionBlocks"
	ClientBoundPlaySelectAdvancementTab            ClientBoundPacket = "SelectAdvancementTab"
	ClientBoundPlayServerData                      ClientBoundPacket = "ServerData"
	ClientBoundPlaySetActionBarText                ClientBoundPacket = "SetActionBarText"
	ClientBoundPlaySetBorderCenter                 ClientBoundPacket = "SetBorderCenter"
	ClientBoundPlaySetBorderLerpSize               ClientBoundPacket = "SetBorderLerpSize"
	ClientBoundPlaySetBorderSize                   ClientBoundPacket = "SetBorderSize"
	ClientBoundPlaySetBorderWarningDelay           ClientBoundPacket = "SetBorderWarningDelay"
	ClientBoundPlaySetBorderWarningDistance        ClientBoundPacket = "SetBorderWarningDistance"
	ClientBoundPlaySetCamera                       ClientBoundPacket = "SetCamera"
	ClientBoundPlaySetHeldItem                     ClientBoundPacket = "SetHeldItem"
	ClientBoundPlaySetCenterChunk                  ClientBoundPacket = "SetCenterChunk"
	ClientBoundPlaySetRenderDistance               ClientBoundPacket = "SetRenderDistance"
	ClientBoundPlaySetDefaultSpawnPosition         ClientBoundPacket = "SetDefaultSpawnPosition"
	ClientBoundPlayDisplayObjective                ClientBoundPacket = "DisplayObjective"
	ClientBoundPlaySetEntityMetadata               ClientBoundPacket = "SetEntityMetadata"
	ClientBoundPlayLinkEntities                    ClientBoundPacket = "LinkEntities"
	ClientBoundPlaySetEntityVelocity               ClientBoundPacket = "SetEntityVelocity"
	ClientBoundPlaySetEquipment                    ClientBoundPacket = "SetEquipment"
	ClientBoundPlaySetExperience                   ClientBoundPacket = "SetExperience"
	ClientBoundPlaySetHealth                       ClientBoundPacket = "SetHealth"
	ClientBoundPlayUpdateObjectives                ClientBoundPacket = "UpdateObjectives"
	ClientBoundPlaySetPassengers                   ClientBoundPacket = "SetPassengers"
	ClientBoundPlayUpdateTeams                     ClientBoundPacket = "UpdateTeams"
	ClientBoundPlayUpdateScore                     ClientBoundPacket = "UpdateScore"
	ClientBoundPlaySetSimulationDistance           ClientBoundPacket = "SetSimulationDistance"
	ClientBoundPlaySetSubtitleText                 ClientBoundPacket = "SetSubtitleText"
	ClientBoundPlayUpdateTime                      ClientBoundPacket = "UpdateTime"
	ClientBoundPlaySetTitleText                    ClientBoundPacket = "SetTitleText"
	ClientBoundPlaySetTitleAnimationTimes          ClientBoundPacket = "SetTitleAnimationTimes"
	ClientBoundPlayEntitySoundEffect               ClientBoundPacket = "EntitySoundEffect"
	ClientBoundPlaySoundEffect                     ClientBoundPacket = "SoundEffect"
	ClientBoundPlayStartConfiguration              ClientBoundPacket = "StartConfiguration"
	ClientBoundPlayStopSound                       ClientBoundPacket = "StopSound"
	ClientBoundPlaySystemChatMessage               ClientBoundPacket = "SystemChatMessage"
	ClientBoundPlaySetTabListHeaderAndFooter       ClientBoundPacket = "SetTabListHeaderAndFooter"
	ClientBoundPlayTagQueryResponse                ClientBoundPacket = "TagQueryResponse"
	ClientBoundPlayPickupItem                      ClientBoundPacket = "PickupItem"
	ClientBoundPlaySetTickingState                 ClientBoundPacket = "SetTickingState"
	ClientBoundPlayStepTick                        ClientBoundPacket = "StepTick"
	ClientBoundPlayTeleportEntity                  ClientBoundPacket = "TeleportEntity"
	ClientBoundPlayUpdateAdvancements              ClientBoundPacket = "UpdateAdvancements"
	ClientBoundPlayUpdateAttributes                ClientBoundPacket = "UpdateAttributes"
	ClientBoundPlayFeatureFlags                    ClientBoundPacket = "FeatureFlags"
	ClientBoundPlayEntityEffect                    ClientBoundPacket = "EntityEffect"
	ClientBoundPlayUpdateRecipes                   ClientBoundPacket = "UpdateRecipes"
	ClientBoundPlayUpdateTags                      ClientBoundPacket = "UpdateTags"
)

type ServerBoundPacket string

const (

	// Handshake
	ServerBoundHandshake                     ServerBoundPacket = "Handshake"
	ServerBoundHandshakeLegacyServerListPing ServerBoundPacket = "LegacyServerListPing"

	// Status
	ServerBoundStatusRequest ServerBoundPacket = "StatusRequest"
	ServerBoundStatusPing    ServerBoundPacket = "StatusPing"

	// Configuration
	ServerBoundConfigurationClientInformation    ServerBoundPacket = "ClientInformation"
	ServerBoundConfigurationPluginMessage        ServerBoundPacket = "PluginMessage"
	ServerBoundConfigurationFinish               ServerBoundPacket = "Finish"
	ServerBoundConfigurationKeepAlive            ServerBoundPacket = "KeepAlive"
	ServerBoundConfigurationPong                 ServerBoundPacket = "Pong"
	ServerBoundConfigurationResourcePackResponse ServerBoundPacket = "ResourcePackResponse"

	// Login
	ServerBoundLoginStart              ServerBoundPacket = "LoginStart"
	ServerBoundLoginEncryptionResponse ServerBoundPacket = "LoginEncryptionResponse"
	ServerBoundLoginPluginResponse     ServerBoundPacket = "LoginPluginResponse"
	ServerBoundLoginAcknowledged       ServerBoundPacket = "LoginAcknowledged"

	// Play
	ServerBoundPlayConfirmTeleportation         ServerBoundPacket = "ConfirmTeleportation"
	ServerBoundPlayQueryBlockEntityTag          ServerBoundPacket = "QueryBlockEntityTag"
	ServerBoundPlayChangeDifficulty             ServerBoundPacket = "ChangeDifficulty"
	ServerBoundPlayMessageAcknowledgement       ServerBoundPacket = "MessageAcknowledgement"
	ServerBoundPlayChatCommand                  ServerBoundPacket = "ChatCommand"
	ServerBoundPlayChatMessage                  ServerBoundPacket = "ChatMessage"
	ServerBoundPlayPlayerSession                ServerBoundPacket = "PlayerSession"
	ServerBoundPlayChunkBatchReceived           ServerBoundPacket = "ChunkBatchRecieved"
	ServerBoundPlayClientCommand                ServerBoundPacket = "ClientCommand"
	ServerBoundPlayClientInformation            ServerBoundPacket = "ClientInformation"
	ServerBoundPlayCommandSuggestionsRequest    ServerBoundPacket = "CommandSuggestionsRequest"
	ServerBoundPlayConfigurationAcknowledged    ServerBoundPacket = "ConfigurationAcknowledged"
	ServerBoundPlayClickContainerButton         ServerBoundPacket = "ClickContainerButton"
	ServerBoundPlayClickContainer               ServerBoundPacket = "ClickContainer"
	ServerBoundPlayCloseContainer               ServerBoundPacket = "CloseContainer"
	ServerBoundPlayChangeContainerSlotState     ServerBoundPacket = "ChangeContainerSlotState"
	ServerBoundPlayPluginMessage                ServerBoundPacket = "PluginMessage"
	ServerBoundPlayEditBook                     ServerBoundPacket = "EditBook"
	ServerBoundPlayQueryEntityTag               ServerBoundPacket = "QueryEntityTag"
	ServerBoundPlayInteractEntity               ServerBoundPacket = "InteractEntity"
	ServerBoundPlayJigsawGenerate               ServerBoundPacket = "JigsawGenerate"
	ServerBoundPlayKeepAlive                    ServerBoundPacket = "KeepAlive"
	ServerBoundPlayLockDifficulty               ServerBoundPacket = "LockDifficulty"
	ServerBoundPlaySetPlayerPosition            ServerBoundPacket = "SetPlayerPosition"
	ServerBoundPlaySetPlayerPositionAndRotation ServerBoundPacket = "SetPlayerPositionAndRotation"
	ServerBoundPlaySetPlayerRotation            ServerBoundPacket = "SetPlayerRotation"
	ServerBoundPlaySetPlayerOnGround            ServerBoundPacket = "SetPlayerOnGround"
	ServerBoundPlayMoveVehicle                  ServerBoundPacket = "MoveVehicle"
	ServerBoundPlayPaddleBoat                   ServerBoundPacket = "PaddleBoat"
	ServerBoundPlayPickItem                     ServerBoundPacket = "PickItem"
	ServerBoundPlayPingRequest                  ServerBoundPacket = "PingRequest"
	ServerBoundPlayPlaceRecipe                  ServerBoundPacket = "PlaceRecipe"
	ServerBoundPlayPlayerAbilities              ServerBoundPacket = "PlayerAbilities"
	ServerBoundPlayPlayerAction                 ServerBoundPacket = "PlayerAction"
	ServerBoundPlayPlayerCommand                ServerBoundPacket = "PlayerCommand"
	ServerBoundPlayPlayerInput                  ServerBoundPacket = "PlayerInput"
	ServerBoundPlayPong                         ServerBoundPacket = "Pong"
	ServerBoundPlayChangeRecipeBookSettings     ServerBoundPacket = "ChangeRecipeBookSettings"
	ServerBoundPlaySetSeenRecipe                ServerBoundPacket = "SetSeenRecipe"
	ServerBoundPlayRenameItem                   ServerBoundPacket = "RenameItem"
	ServerBoundPlayResourcePack                 ServerBoundPacket = "ResourcePack"
	ServerBoundPlaySeenAdvancements             ServerBoundPacket = "SeenAdvancements"
	ServerBoundPlaySelectTrade                  ServerBoundPacket = "SelectTrade"
	ServerBoundPlaySetBeaconEffect              ServerBoundPacket = "SetBeaconEffect"
	ServerBoundPlaySetHeldItem                  ServerBoundPacket = "SetHeldItem"
	ServerBoundPlayProgramCommandBlock          ServerBoundPacket = "ProgramCommandBlock"
	ServerBoundPlayProgramCommandBlockMinecart  ServerBoundPacket = "ProgramCommandBlockMinecart"
	ServerBoundPlaySetCreativeModeSlot          ServerBoundPacket = "SetCreativeModeSlot"
	ServerBoundPlayProgramJigsawBlock           ServerBoundPacket = "ProgramJigsawBlock"
	ServerBoundPlayProgramStructureBlock        ServerBoundPacket = "ProgramStructureBlock"
	ServerBoundPlayUpdateSign                   ServerBoundPacket = "UpdateSign"
	ServerBoundPlaySwingArm                     ServerBoundPacket = "SwingArm"
	ServerBoundPlayTeleportToEntity             ServerBoundPacket = "TeleportToEntity"
	ServerBoundPlayUseItemOn                    ServerBoundPacket = "UseItemOn"
	ServerBoundPlayUseItem                      ServerBoundPacket = "UseItem"
)
