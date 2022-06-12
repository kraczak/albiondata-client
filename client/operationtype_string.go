// Code generated by "stringer -type=OperationType"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opUnused-0]
	_ = x[opPing-1]
	_ = x[opJoin-2]
	_ = x[opCreateAccount-3]
	_ = x[opLogin-4]
	_ = x[opSendCrashLog-5]
	_ = x[opSendTraceRoute-6]
	_ = x[opSendVfxStats-7]
	_ = x[opSendGamePingInfo-8]
	_ = x[opCreateCharacter-9]
	_ = x[opDeleteCharacter-10]
	_ = x[opSelectCharacter-11]
	_ = x[opRedeemKeycode-12]
	_ = x[opGetGameServerByCluster-13]
	_ = x[opGetActiveSubscription-14]
	_ = x[opGetShopPurchaseUrl-15]
	_ = x[opGetBuyTrialDetails-16]
	_ = x[opGetReferralSeasonDetails-17]
	_ = x[opGetReferralLink-18]
	_ = x[opGetAvailableTrialKeys-19]
	_ = x[opGetShopTilesForCategory-20]
	_ = x[opMove-21]
	_ = x[opCastStart-22]
	_ = x[opCastCancel-23]
	_ = x[opTerminateToggleSpell-24]
	_ = x[opChannelingCancel-25]
	_ = x[opAttackBuildingStart-26]
	_ = x[opInventoryDestroyItem-27]
	_ = x[opInventoryMoveItem-28]
	_ = x[opInventoryRecoverItem-29]
	_ = x[opInventoryRecoverAllItems-30]
	_ = x[opInventorySplitStack-31]
	_ = x[opInventorySplitStackInto-32]
	_ = x[opGetClusterData-33]
	_ = x[opChangeCluster-34]
	_ = x[opConsoleCommand-35]
	_ = x[opChatMessage-36]
	_ = x[opReportClientError-37]
	_ = x[opRegisterToObject-38]
	_ = x[opUnRegisterFromObject-39]
	_ = x[opCraftBuildingChangeSettings-40]
	_ = x[opCraftBuildingTakeMoney-41]
	_ = x[opRepairBuildingChangeSettings-42]
	_ = x[opRepairBuildingTakeMoney-43]
	_ = x[opActionBuildingChangeSettings-44]
	_ = x[opHarvestStart-45]
	_ = x[opHarvestCancel-46]
	_ = x[opTakeSilver-47]
	_ = x[opActionOnBuildingStart-48]
	_ = x[opActionOnBuildingCancel-49]
	_ = x[opItemRerollQualityStart-50]
	_ = x[opItemRerollQualityCancel-51]
	_ = x[opInstallResourceStart-52]
	_ = x[opInstallResourceCancel-53]
	_ = x[opInstallSilver-54]
	_ = x[opBuildingFillNutrition-55]
	_ = x[opBuildingChangeRenovationState-56]
	_ = x[opBuildingBuySkin-57]
	_ = x[opBuildingClaim-58]
	_ = x[opBuildingGiveup-59]
	_ = x[opBuildingNutritionSilverStorageDeposit-60]
	_ = x[opBuildingNutritionSilverStorageWithdraw-61]
	_ = x[opBuildingNutritionSilverRewardSet-62]
	_ = x[opConstructionSiteCreate-63]
	_ = x[opPlaceableObjectPlace-64]
	_ = x[opPlaceableObjectPlaceCancel-65]
	_ = x[opPlaceableObjectPickup-66]
	_ = x[opFurnitureObjectUse-67]
	_ = x[opFarmableHarvest-68]
	_ = x[opFarmableFinishGrownItem-69]
	_ = x[opFarmableDestroy-70]
	_ = x[opFarmableGetProduct-71]
	_ = x[opTearDownConstructionSite-72]
	_ = x[opCastleGateUse-73]
	_ = x[opAuctionCreateRequest-74]
	_ = x[opAuctionCreateOffer-75]
	_ = x[opAuctionGetOffers-76]
	_ = x[opAuctionGetRequests-77]
	_ = x[opAuctionBuyOffer-78]
	_ = x[opAuctionAbortAuction-79]
	_ = x[opAuctionModifyAuction-80]
	_ = x[opAuctionAbortOffer-81]
	_ = x[opAuctionAbortRequest-82]
	_ = x[opAuctionSellRequest-83]
	_ = x[opAuctionGetFinishedAuctions-84]
	_ = x[opAuctionGetFinishedAuctionsCount-85]
	_ = x[opAuctionFetchAuction-86]
	_ = x[opAuctionGetMyOpenOffers-87]
	_ = x[opAuctionGetMyOpenRequests-88]
	_ = x[opAuctionGetMyOpenAuctions-89]
	_ = x[opUnknown90-90]
	_ = x[opAuctionGetItemAverageStats-91]
	_ = x[opAuctionGetItemAverageValue-92]
	_ = x[opContainerOpen-93]
	_ = x[opContainerClose-94]
	_ = x[opContainerManageSubContainer-95]
	_ = x[opRespawn-96]
	_ = x[opSuicide-97]
	_ = x[opJoinGuild-98]
	_ = x[opLeaveGuild-99]
	_ = x[opCreateGuild-100]
	_ = x[opInviteToGuild-101]
	_ = x[opDeclineGuildInvitation-102]
	_ = x[opKickFromGuild-103]
	_ = x[opDuellingChallengePlayer-104]
	_ = x[opDuellingAcceptChallenge-105]
	_ = x[opDuellingDenyChallenge-106]
	_ = x[opChangeClusterTax-107]
	_ = x[opClaimTerritory-108]
	_ = x[opGiveUpTerritory-109]
	_ = x[opChangeTerritoryAccessRights-110]
	_ = x[opGetMonolithInfo-111]
	_ = x[opGetClaimInfo-112]
	_ = x[opGetAttackInfo-113]
	_ = x[opGetTerritorySeasonPoints-114]
	_ = x[opGetAttackSchedule-115]
	_ = x[opScheduleAttack-116]
	_ = x[opGetMatches-117]
	_ = x[opGetMatchDetails-118]
	_ = x[opJoinMatch-119]
	_ = x[opLeaveMatch-120]
	_ = x[opChangeChatSettings-121]
	_ = x[opLogoutStart-122]
	_ = x[opLogoutCancel-123]
	_ = x[opClaimOrbStart-124]
	_ = x[opClaimOrbCancel-125]
	_ = x[opMatchLootChestOpeningStart-126]
	_ = x[opMatchLootChestOpeningCancel-127]
	_ = x[opDepositToGuildAccount-128]
	_ = x[opWithdrawalFromAccount-129]
	_ = x[opChangeGuildPayUpkeepFlag-130]
	_ = x[opChangeGuildTax-131]
	_ = x[opGetMyTerritories-132]
	_ = x[opMorganaCommand-133]
	_ = x[opGetServerInfo-134]
	_ = x[opInviteMercenaryToMatch-135]
	_ = x[opSubscribeToCluster-136]
	_ = x[opAnswerMercenaryInvitation-137]
	_ = x[opGetCharacterEquipment-138]
	_ = x[opGetCharacterSteamAchievements-139]
	_ = x[opGetCharacterStats-140]
	_ = x[opGetKillHistoryDetails-141]
	_ = x[opLearnMasteryLevel-142]
	_ = x[opReSpecAchievement-143]
	_ = x[opChangeAvatar-144]
	_ = x[opGetRankings-145]
	_ = x[opGetRank-146]
	_ = x[opGetGvgSeasonRankings-147]
	_ = x[opGetGvgSeasonRank-148]
	_ = x[opGetGvgSeasonHistoryRankings-149]
	_ = x[opGetGvgSeasonGuildMemberHistory-150]
	_ = x[opKickFromGvGMatch-151]
	_ = x[opGetChestLogs-152]
	_ = x[opGetAccessRightLogs-153]
	_ = x[opGetGuildAccountLogs-154]
	_ = x[opGetGuildAccountLogsLargeAmount-155]
	_ = x[opInviteToPlayerTrade-156]
	_ = x[opPlayerTradeCancel-157]
	_ = x[opPlayerTradeInvitationAccept-158]
	_ = x[opPlayerTradeAddItem-159]
	_ = x[opPlayerTradeRemoveItem-160]
	_ = x[opPlayerTradeAcceptTrade-161]
	_ = x[opPlayerTradeSetSilverOrGold-162]
	_ = x[opSendMiniMapPing-163]
	_ = x[opStuck-164]
	_ = x[opBuyRealEstate-165]
	_ = x[opClaimRealEstate-166]
	_ = x[opGiveUpRealEstate-167]
	_ = x[opChangeRealEstateOutline-168]
	_ = x[opGetMailInfos-169]
	_ = x[opGetMailCount-170]
	_ = x[opReadMail-171]
	_ = x[opSendNewMail-172]
	_ = x[opDeleteMail-173]
	_ = x[opMarkMailUnread-174]
	_ = x[opClaimAttachmentFromMail-175]
	_ = x[opUpdateLfgInfo-176]
	_ = x[opGetLfgInfos-177]
	_ = x[opGetMyGuildLfgInfo-178]
	_ = x[opGetLfgDescriptionText-179]
	_ = x[opLfgApplyToGuild-180]
	_ = x[opAnswerLfgGuildApplication-181]
	_ = x[opRegisterChatPeer-182]
	_ = x[opSendChatMessage-183]
	_ = x[opJoinChatChannel-184]
	_ = x[opLeaveChatChannel-185]
	_ = x[opSendWhisperMessage-186]
	_ = x[opSay-187]
	_ = x[opPlayEmote-188]
	_ = x[opStopEmote-189]
	_ = x[opGetClusterMapInfo-190]
	_ = x[opAccessRightsChangeSettings-191]
	_ = x[opMount-192]
	_ = x[opMountCancel-193]
	_ = x[opBuyJourney-194]
	_ = x[opSetSaleStatusForEstate-195]
	_ = x[opResolveGuildOrPlayerName-196]
	_ = x[opGetRespawnInfos-197]
	_ = x[opMakeHome-198]
	_ = x[opLeaveHome-199]
	_ = x[opResurrectionReply-200]
	_ = x[opAllianceCreate-201]
	_ = x[opAllianceDisband-202]
	_ = x[opAllianceGetMemberInfos-203]
	_ = x[opAllianceInvite-204]
	_ = x[opAllianceAnswerInvitation-205]
	_ = x[opAllianceCancelInvitation-206]
	_ = x[opAllianceKickGuild-207]
	_ = x[opAllianceLeave-208]
	_ = x[opAllianceChangeGoldPaymentFlag-209]
	_ = x[opAllianceGetDetailInfo-210]
	_ = x[opGetIslandInfos-211]
	_ = x[opAbandonMyIsland-212]
	_ = x[opBuyMyIsland-213]
	_ = x[opBuyGuildIsland-214]
	_ = x[opAbandonGuildIsland-215]
	_ = x[opUpgradeMyIsland-216]
	_ = x[opUpgradeGuildIsland-217]
	_ = x[opMoveMyIsland-218]
	_ = x[opMoveGuildIsland-219]
	_ = x[opTerritoryFillNutrition-220]
	_ = x[opTeleportBack-221]
	_ = x[opPartyInvitePlayer-222]
	_ = x[opPartyAnswerInvitation-223]
	_ = x[opPartyLeave-224]
	_ = x[opPartyKickPlayer-225]
	_ = x[opPartyMakeLeader-226]
	_ = x[opPartyChangeLootSetting-227]
	_ = x[opPartyMarkObject-228]
	_ = x[opPartySetRole-229]
	_ = x[opGetGuildMOTD-230]
	_ = x[opSetGuildMOTD-231]
	_ = x[opExitEnterStart-232]
	_ = x[opExitEnterCancel-233]
	_ = x[opQuestGiverRequest-234]
	_ = x[opGoldMarketGetBuyOffer-235]
	_ = x[opGoldMarketGetBuyOfferFromSilver-236]
	_ = x[opGoldMarketGetSellOffer-237]
	_ = x[opGoldMarketGetSellOfferFromSilver-238]
	_ = x[opGoldMarketBuyGold-239]
	_ = x[opGoldMarketSellGold-240]
	_ = x[opGoldMarketCreateSellOrder-241]
	_ = x[opGoldMarketCreateBuyOrder-242]
	_ = x[opGoldMarketGetInfos-243]
	_ = x[opGoldMarketCancelOrder-244]
	_ = x[opUnknown244-245]
	_ = x[opUnknown245-246]
	_ = x[opGoldMarketGetAverageInfo-247]
	_ = x[opSiegeCampClaimStart-248]
	_ = x[opSiegeCampClaimCancel-249]
	_ = x[opTreasureChestUsingStart-250]
	_ = x[opTreasureChestUsingCancel-251]
	_ = x[opUseLootChest-252]
	_ = x[opUseShrine-253]
	_ = x[opLaborerStartJob-254]
	_ = x[opLaborerTakeJobLoot-255]
	_ = x[opLaborerDismiss-256]
	_ = x[opLaborerMove-257]
	_ = x[opLaborerBuyItem-258]
	_ = x[opLaborerUpgrade-259]
	_ = x[opBuyPremium-260]
	_ = x[opBuyTrial-261]
	_ = x[opRealEstateGetAuctionData-262]
	_ = x[opRealEstateBidOnAuction-263]
	_ = x[opGetSiegeCampCooldown-264]
	_ = x[opFriendInvite-265]
	_ = x[opFriendAnswerInvitation-266]
	_ = x[opFriendCancelnvitation-267]
	_ = x[opFriendRemove-268]
	_ = x[opInventoryStack-269]
	_ = x[opInventorySort-270]
	_ = x[opEquipmentItemChangeSpell-271]
	_ = x[opExpeditionRegister-272]
	_ = x[opExpeditionRegisterCancel-273]
	_ = x[opJoinExpedition-274]
	_ = x[opDeclineExpeditionInvitation-275]
	_ = x[opVoteStart-276]
	_ = x[opVoteDoVote-277]
	_ = x[opRatingDoRate-278]
	_ = x[opEnteringExpeditionStart-279]
	_ = x[opEnteringExpeditionCancel-280]
	_ = x[opActivateExpeditionCheckPoint-281]
	_ = x[opArenaRegister-282]
	_ = x[opArenaRegisterCancel-283]
	_ = x[opArenaLeave-284]
	_ = x[opJoinArenaMatch-285]
	_ = x[opDeclineArenaInvitation-286]
	_ = x[opEnteringArenaStart-287]
	_ = x[opEnteringArenaCancel-288]
	_ = x[opArenaCustomMatch-289]
	_ = x[opArenaCustomMatchCreate-290]
	_ = x[opUpdateCharacterStatement-291]
	_ = x[opBoostFarmable-292]
	_ = x[opGetStrikeHistory-293]
	_ = x[opUseFunction-294]
	_ = x[opUsePortalEntrance-295]
	_ = x[opResetPortalBinding-296]
	_ = x[opQueryPortalBinding-297]
	_ = x[opClaimPaymentTransaction-298]
	_ = x[opChangeUseFlag-299]
	_ = x[opClientPerformanceStats-300]
	_ = x[opExtendedHardwareStats-301]
	_ = x[opClientLowMemoryWarning-302]
	_ = x[opTerritoryClaimStart-303]
	_ = x[opTerritoryClaimCancel-304]
	_ = x[opRequestAppStoreProducts-305]
	_ = x[opVerifyProductPurchase-306]
	_ = x[opQueryGuildPlayerStats-307]
	_ = x[opQueryAllianceGuildStats-308]
	_ = x[opTrackAchievements-309]
	_ = x[opSetAchievementsAutoLearn-310]
	_ = x[opDepositItemToGuildCurrency-311]
	_ = x[opWithdrawalItemFromGuildCurrency-312]
	_ = x[opAuctionSellSpecificItemRequest-313]
	_ = x[opFishingStart-314]
	_ = x[opFishingCasting-315]
	_ = x[opFishingCast-316]
	_ = x[opFishingCatch-317]
	_ = x[opFishingPull-318]
	_ = x[opFishingGiveLine-319]
	_ = x[opFishingFinish-320]
	_ = x[opFishingCancel-321]
	_ = x[opCreateGuildAccessTag-322]
	_ = x[opDeleteGuildAccessTag-323]
	_ = x[opRenameGuildAccessTag-324]
	_ = x[opFlagGuildAccessTagGuildPermission-325]
	_ = x[opAssignGuildAccessTag-326]
	_ = x[opRemoveGuildAccessTagFromPlayer-327]
	_ = x[opModifyGuildAccessTagEditors-328]
	_ = x[opRequestPublicAccessTags-329]
	_ = x[opChangeAccessTagPublicFlag-330]
	_ = x[opUpdateGuildAccessTag-331]
	_ = x[opSteamStartMicrotransaction-332]
	_ = x[opSteamFinishMicrotransaction-333]
	_ = x[opSteamIdHasActiveAccount-334]
	_ = x[opCheckEmailAccountState-335]
	_ = x[opLinkAccountToSteamId-336]
	_ = x[opBuyGvgSeasonBooster-337]
	_ = x[opChangeFlaggingPrepare-338]
	_ = x[opOverCharge-339]
	_ = x[opOverChargeEnd-340]
	_ = x[opRequestTrusted-341]
	_ = x[opChangeGuildLogo-342]
	_ = x[opPartyFinderRegisterForUpdates-343]
	_ = x[opPartyFinderUnregisterForUpdates-344]
	_ = x[opPartyFinderEnlistNewPartySearch-345]
	_ = x[opPartyFinderDeletePartySearch-346]
	_ = x[opPartyFinderChangePartySearch-347]
	_ = x[opPartyFinderChangeRole-348]
	_ = x[opPartyFinderApplyForGroup-349]
	_ = x[opPartyFinderAcceptOrDeclineApplyForGroup-350]
	_ = x[opPartyFinderGetEquipmentSnapshot-351]
	_ = x[opPartyFinderRegisterApplicants-352]
	_ = x[opPartyFinderUnregisterApplicants-353]
	_ = x[opPartyFinderFulltextSearch-354]
	_ = x[opPartyFinderRequestEquipmentSnapshot-355]
	_ = x[opGetPersonalSeasonTrackerData-356]
	_ = x[opUseConsumableFromInventory-357]
	_ = x[opClaimPersonalSeasonReward-358]
	_ = x[opEasyAntiCheatMessageToServer-359]
	_ = x[opSetNextTutorialState-360]
	_ = x[opAddPlayerToMuteList-361]
	_ = x[opRemovePlayerFromMuteList-362]
	_ = x[opProductShopUserEvent-363]
	_ = x[opGetVanityUnlocks-364]
	_ = x[opBuyVanityUnlocks-365]
	_ = x[opGetMountSkins-366]
	_ = x[opSetMountSkin-367]
	_ = x[opSetWardrobe-368]
	_ = x[opChangeCustomization-369]
	_ = x[opSetFavoriteIsland-370]
	_ = x[opGetGuildChallengePoints-371]
	_ = x[opTravelToHideout-372]
	_ = x[opSmartQueueJoin-373]
	_ = x[opSmartQueueLeave-374]
	_ = x[opSmartQueueSelectSpawnCluster-375]
	_ = x[opUpgradeHideout-376]
	_ = x[opInitHideoutAttackStart-377]
	_ = x[opInitHideoutAttackCancel-378]
	_ = x[opHideoutFillNutrition-379]
	_ = x[opHideoutGetInfo-380]
	_ = x[opHideoutGetOwnerInfo-381]
	_ = x[opHideoutSetTribute-382]
	_ = x[opOpenWorldAttackScheduleStart-383]
	_ = x[opOpenWorldAttackScheduleCancel-384]
	_ = x[opOpenWorldAttackConquerStart-385]
	_ = x[opOpenWorldAttackConquerCancel-386]
	_ = x[opGetOpenWorldAttackDetails-387]
	_ = x[opGetNextOpenWorldAttackScheduleTime-388]
	_ = x[opRecoverVaultFromHideout-389]
	_ = x[opGetGuildEnergyDrainInfo-390]
	_ = x[opChannelingUpdate-391]
}

const _OperationType_name = "opUnusedopPingopJoinopCreateAccountopLoginopSendCrashLogopSendTraceRouteopSendVfxStatsopSendGamePingInfoopCreateCharacteropDeleteCharacteropSelectCharacteropRedeemKeycodeopGetGameServerByClusteropGetActiveSubscriptionopGetShopPurchaseUrlopGetBuyTrialDetailsopGetReferralSeasonDetailsopGetReferralLinkopGetAvailableTrialKeysopGetShopTilesForCategoryopMoveopCastStartopCastCancelopTerminateToggleSpellopChannelingCancelopAttackBuildingStartopInventoryDestroyItemopInventoryMoveItemopInventoryRecoverItemopInventoryRecoverAllItemsopInventorySplitStackopInventorySplitStackIntoopGetClusterDataopChangeClusteropConsoleCommandopChatMessageopReportClientErroropRegisterToObjectopUnRegisterFromObjectopCraftBuildingChangeSettingsopCraftBuildingTakeMoneyopRepairBuildingChangeSettingsopRepairBuildingTakeMoneyopActionBuildingChangeSettingsopHarvestStartopHarvestCancelopTakeSilveropActionOnBuildingStartopActionOnBuildingCancelopItemRerollQualityStartopItemRerollQualityCancelopInstallResourceStartopInstallResourceCancelopInstallSilveropBuildingFillNutritionopBuildingChangeRenovationStateopBuildingBuySkinopBuildingClaimopBuildingGiveupopBuildingNutritionSilverStorageDepositopBuildingNutritionSilverStorageWithdrawopBuildingNutritionSilverRewardSetopConstructionSiteCreateopPlaceableObjectPlaceopPlaceableObjectPlaceCancelopPlaceableObjectPickupopFurnitureObjectUseopFarmableHarvestopFarmableFinishGrownItemopFarmableDestroyopFarmableGetProductopTearDownConstructionSiteopCastleGateUseopAuctionCreateRequestopAuctionCreateOfferopAuctionGetOffersopAuctionGetRequestsopAuctionBuyOfferopAuctionAbortAuctionopAuctionModifyAuctionopAuctionAbortOfferopAuctionAbortRequestopAuctionSellRequestopAuctionGetFinishedAuctionsopAuctionGetFinishedAuctionsCountopAuctionFetchAuctionopAuctionGetMyOpenOffersopAuctionGetMyOpenRequestsopAuctionGetMyOpenAuctionsopUnknown90opAuctionGetItemAverageStatsopAuctionGetItemAverageValueopContainerOpenopContainerCloseopContainerManageSubContaineropRespawnopSuicideopJoinGuildopLeaveGuildopCreateGuildopInviteToGuildopDeclineGuildInvitationopKickFromGuildopDuellingChallengePlayeropDuellingAcceptChallengeopDuellingDenyChallengeopChangeClusterTaxopClaimTerritoryopGiveUpTerritoryopChangeTerritoryAccessRightsopGetMonolithInfoopGetClaimInfoopGetAttackInfoopGetTerritorySeasonPointsopGetAttackScheduleopScheduleAttackopGetMatchesopGetMatchDetailsopJoinMatchopLeaveMatchopChangeChatSettingsopLogoutStartopLogoutCancelopClaimOrbStartopClaimOrbCancelopMatchLootChestOpeningStartopMatchLootChestOpeningCancelopDepositToGuildAccountopWithdrawalFromAccountopChangeGuildPayUpkeepFlagopChangeGuildTaxopGetMyTerritoriesopMorganaCommandopGetServerInfoopInviteMercenaryToMatchopSubscribeToClusteropAnswerMercenaryInvitationopGetCharacterEquipmentopGetCharacterSteamAchievementsopGetCharacterStatsopGetKillHistoryDetailsopLearnMasteryLevelopReSpecAchievementopChangeAvataropGetRankingsopGetRankopGetGvgSeasonRankingsopGetGvgSeasonRankopGetGvgSeasonHistoryRankingsopGetGvgSeasonGuildMemberHistoryopKickFromGvGMatchopGetChestLogsopGetAccessRightLogsopGetGuildAccountLogsopGetGuildAccountLogsLargeAmountopInviteToPlayerTradeopPlayerTradeCancelopPlayerTradeInvitationAcceptopPlayerTradeAddItemopPlayerTradeRemoveItemopPlayerTradeAcceptTradeopPlayerTradeSetSilverOrGoldopSendMiniMapPingopStuckopBuyRealEstateopClaimRealEstateopGiveUpRealEstateopChangeRealEstateOutlineopGetMailInfosopGetMailCountopReadMailopSendNewMailopDeleteMailopMarkMailUnreadopClaimAttachmentFromMailopUpdateLfgInfoopGetLfgInfosopGetMyGuildLfgInfoopGetLfgDescriptionTextopLfgApplyToGuildopAnswerLfgGuildApplicationopRegisterChatPeeropSendChatMessageopJoinChatChannelopLeaveChatChannelopSendWhisperMessageopSayopPlayEmoteopStopEmoteopGetClusterMapInfoopAccessRightsChangeSettingsopMountopMountCancelopBuyJourneyopSetSaleStatusForEstateopResolveGuildOrPlayerNameopGetRespawnInfosopMakeHomeopLeaveHomeopResurrectionReplyopAllianceCreateopAllianceDisbandopAllianceGetMemberInfosopAllianceInviteopAllianceAnswerInvitationopAllianceCancelInvitationopAllianceKickGuildopAllianceLeaveopAllianceChangeGoldPaymentFlagopAllianceGetDetailInfoopGetIslandInfosopAbandonMyIslandopBuyMyIslandopBuyGuildIslandopAbandonGuildIslandopUpgradeMyIslandopUpgradeGuildIslandopMoveMyIslandopMoveGuildIslandopTerritoryFillNutritionopTeleportBackopPartyInvitePlayeropPartyAnswerInvitationopPartyLeaveopPartyKickPlayeropPartyMakeLeaderopPartyChangeLootSettingopPartyMarkObjectopPartySetRoleopGetGuildMOTDopSetGuildMOTDopExitEnterStartopExitEnterCancelopQuestGiverRequestopGoldMarketGetBuyOfferopGoldMarketGetBuyOfferFromSilveropGoldMarketGetSellOfferopGoldMarketGetSellOfferFromSilveropGoldMarketBuyGoldopGoldMarketSellGoldopGoldMarketCreateSellOrderopGoldMarketCreateBuyOrderopGoldMarketGetInfosopGoldMarketCancelOrderopUnknown244opUnknown245opGoldMarketGetAverageInfoopSiegeCampClaimStartopSiegeCampClaimCancelopTreasureChestUsingStartopTreasureChestUsingCancelopUseLootChestopUseShrineopLaborerStartJobopLaborerTakeJobLootopLaborerDismissopLaborerMoveopLaborerBuyItemopLaborerUpgradeopBuyPremiumopBuyTrialopRealEstateGetAuctionDataopRealEstateBidOnAuctionopGetSiegeCampCooldownopFriendInviteopFriendAnswerInvitationopFriendCancelnvitationopFriendRemoveopInventoryStackopInventorySortopEquipmentItemChangeSpellopExpeditionRegisteropExpeditionRegisterCancelopJoinExpeditionopDeclineExpeditionInvitationopVoteStartopVoteDoVoteopRatingDoRateopEnteringExpeditionStartopEnteringExpeditionCancelopActivateExpeditionCheckPointopArenaRegisteropArenaRegisterCancelopArenaLeaveopJoinArenaMatchopDeclineArenaInvitationopEnteringArenaStartopEnteringArenaCancelopArenaCustomMatchopArenaCustomMatchCreateopUpdateCharacterStatementopBoostFarmableopGetStrikeHistoryopUseFunctionopUsePortalEntranceopResetPortalBindingopQueryPortalBindingopClaimPaymentTransactionopChangeUseFlagopClientPerformanceStatsopExtendedHardwareStatsopClientLowMemoryWarningopTerritoryClaimStartopTerritoryClaimCancelopRequestAppStoreProductsopVerifyProductPurchaseopQueryGuildPlayerStatsopQueryAllianceGuildStatsopTrackAchievementsopSetAchievementsAutoLearnopDepositItemToGuildCurrencyopWithdrawalItemFromGuildCurrencyopAuctionSellSpecificItemRequestopFishingStartopFishingCastingopFishingCastopFishingCatchopFishingPullopFishingGiveLineopFishingFinishopFishingCancelopCreateGuildAccessTagopDeleteGuildAccessTagopRenameGuildAccessTagopFlagGuildAccessTagGuildPermissionopAssignGuildAccessTagopRemoveGuildAccessTagFromPlayeropModifyGuildAccessTagEditorsopRequestPublicAccessTagsopChangeAccessTagPublicFlagopUpdateGuildAccessTagopSteamStartMicrotransactionopSteamFinishMicrotransactionopSteamIdHasActiveAccountopCheckEmailAccountStateopLinkAccountToSteamIdopBuyGvgSeasonBoosteropChangeFlaggingPrepareopOverChargeopOverChargeEndopRequestTrustedopChangeGuildLogoopPartyFinderRegisterForUpdatesopPartyFinderUnregisterForUpdatesopPartyFinderEnlistNewPartySearchopPartyFinderDeletePartySearchopPartyFinderChangePartySearchopPartyFinderChangeRoleopPartyFinderApplyForGroupopPartyFinderAcceptOrDeclineApplyForGroupopPartyFinderGetEquipmentSnapshotopPartyFinderRegisterApplicantsopPartyFinderUnregisterApplicantsopPartyFinderFulltextSearchopPartyFinderRequestEquipmentSnapshotopGetPersonalSeasonTrackerDataopUseConsumableFromInventoryopClaimPersonalSeasonRewardopEasyAntiCheatMessageToServeropSetNextTutorialStateopAddPlayerToMuteListopRemovePlayerFromMuteListopProductShopUserEventopGetVanityUnlocksopBuyVanityUnlocksopGetMountSkinsopSetMountSkinopSetWardrobeopChangeCustomizationopSetFavoriteIslandopGetGuildChallengePointsopTravelToHideoutopSmartQueueJoinopSmartQueueLeaveopSmartQueueSelectSpawnClusteropUpgradeHideoutopInitHideoutAttackStartopInitHideoutAttackCancelopHideoutFillNutritionopHideoutGetInfoopHideoutGetOwnerInfoopHideoutSetTributeopOpenWorldAttackScheduleStartopOpenWorldAttackScheduleCancelopOpenWorldAttackConquerStartopOpenWorldAttackConquerCancelopGetOpenWorldAttackDetailsopGetNextOpenWorldAttackScheduleTimeopRecoverVaultFromHideoutopGetGuildEnergyDrainInfoopChannelingUpdate"

var _OperationType_index = [...]uint16{0, 8, 14, 20, 35, 42, 56, 72, 86, 104, 121, 138, 155, 170, 194, 217, 237, 257, 283, 300, 323, 348, 354, 365, 377, 399, 417, 438, 460, 479, 501, 527, 548, 573, 589, 604, 620, 633, 652, 670, 692, 721, 745, 775, 800, 830, 844, 859, 871, 894, 918, 942, 967, 989, 1012, 1027, 1050, 1081, 1098, 1113, 1129, 1168, 1208, 1242, 1266, 1288, 1316, 1339, 1359, 1376, 1401, 1418, 1438, 1464, 1479, 1501, 1521, 1539, 1559, 1576, 1597, 1619, 1638, 1659, 1679, 1707, 1740, 1761, 1785, 1811, 1837, 1848, 1876, 1904, 1919, 1935, 1964, 1973, 1982, 1993, 2005, 2018, 2033, 2057, 2072, 2097, 2122, 2145, 2163, 2179, 2196, 2225, 2242, 2256, 2271, 2297, 2316, 2332, 2344, 2361, 2372, 2384, 2404, 2417, 2431, 2446, 2462, 2490, 2519, 2542, 2565, 2591, 2607, 2625, 2641, 2656, 2680, 2700, 2727, 2750, 2781, 2800, 2823, 2842, 2861, 2875, 2888, 2897, 2919, 2937, 2966, 2998, 3016, 3030, 3050, 3071, 3103, 3124, 3143, 3172, 3192, 3215, 3239, 3267, 3284, 3291, 3306, 3323, 3341, 3366, 3380, 3394, 3404, 3417, 3429, 3445, 3470, 3485, 3498, 3517, 3540, 3557, 3584, 3602, 3619, 3636, 3654, 3674, 3679, 3690, 3701, 3720, 3748, 3755, 3768, 3780, 3804, 3830, 3847, 3857, 3868, 3887, 3903, 3920, 3944, 3960, 3986, 4012, 4031, 4046, 4077, 4100, 4116, 4133, 4146, 4162, 4182, 4199, 4219, 4233, 4250, 4274, 4288, 4307, 4330, 4342, 4359, 4376, 4400, 4417, 4431, 4445, 4459, 4475, 4492, 4511, 4534, 4567, 4591, 4625, 4644, 4664, 4691, 4717, 4737, 4760, 4772, 4784, 4810, 4831, 4853, 4878, 4904, 4918, 4929, 4946, 4966, 4982, 4995, 5011, 5027, 5039, 5049, 5075, 5099, 5121, 5135, 5159, 5182, 5196, 5212, 5227, 5253, 5273, 5299, 5315, 5344, 5355, 5367, 5381, 5406, 5432, 5462, 5477, 5498, 5510, 5526, 5550, 5570, 5591, 5609, 5633, 5659, 5674, 5692, 5705, 5724, 5744, 5764, 5789, 5804, 5828, 5851, 5875, 5896, 5918, 5943, 5966, 5989, 6014, 6033, 6059, 6087, 6120, 6152, 6166, 6182, 6195, 6209, 6222, 6239, 6254, 6269, 6291, 6313, 6335, 6370, 6392, 6424, 6453, 6478, 6505, 6527, 6555, 6584, 6609, 6633, 6655, 6676, 6699, 6711, 6726, 6742, 6759, 6790, 6823, 6856, 6886, 6916, 6939, 6965, 7006, 7039, 7070, 7103, 7130, 7167, 7197, 7225, 7252, 7282, 7304, 7325, 7351, 7373, 7391, 7409, 7424, 7438, 7451, 7472, 7491, 7516, 7533, 7549, 7566, 7596, 7612, 7636, 7661, 7683, 7699, 7720, 7739, 7769, 7800, 7829, 7859, 7886, 7922, 7947, 7972, 7990}

func (i OperationType) String() string {
	if i >= OperationType(len(_OperationType_index)-1) {
		return "OperationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationType_name[_OperationType_index[i]:_OperationType_index[i+1]]
}
