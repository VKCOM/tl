// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"sort"
)

// All compatibility code that will be removed after some fixes to TL itself

type legacyPair struct {
	conName   string
	fieldName string
}

var legacyExceptions []legacyPair // use it to populate this file with content of combined tl

func LegacyPrintGlobalMap() {
	sort.Slice(legacyExceptions, func(i, j int) bool {
		if legacyExceptions[i].conName != legacyExceptions[j].conName {
			return legacyExceptions[i].conName < legacyExceptions[j].conName
		}
		return legacyExceptions[i].fieldName < legacyExceptions[j].fieldName
	})

	for _, v := range legacyExceptions {
		fmt.Printf(`	if conFullName == %q && fieldName == %q {
		return true
	}
`, v.conName, v.fieldName)
	}
	if len(legacyExceptions) != 0 {
		fmt.Printf("--- Total %d exceptions\n", len(legacyExceptions))
	}
}

func LegacyEnableExclamation(conFullName string) bool {
	if conFullName == "rpcDestActor" {
		return true
	}
	if conFullName == "rpcDestActorFlags" {
		return true
	}
	if conFullName == "rpcDestFlags" {
		return true
	}
	if conFullName == "rpcInvokeReq" {
		return true
	}
	if conFullName == "engine.sendResponseTo" {
		return true
	}
	if conFullName == "messages.responseQuery" {
		return true
	}
	if conFullName == "rpcProxy.diagonal" {
		return true
	}
	if conFullName == "rpcProxy.diagonalTargets" {
		return true
	}
	if conFullName == "storage2Impl.forwardToFork" {
		return true
	}
	if conFullName == "storage2Impl.diagonalFork" {
		return true
	}
	if conFullName == "tree_stats.preferMaster" {
		return true
	}
	return false
}

func LegacyEnableWarningsUnionNamespaceSkip(conNamespace string, typeNamespace string) bool {
	if conNamespace == "messagesLong" && typeNamespace == "messages" {
		return true
	}
	if conNamespace == "storage2Impl" && typeNamespace == "storage2" {
		return true
	}
	if conNamespace == "expr" && typeNamespace == "" {
		return true
	}
	if conNamespace == "gucene" && typeNamespace == "" {
		return true
	}
	// below is goldmaster, we keep it to test generated code for this case
	if conNamespace == "cd" && typeNamespace == "ab" {
		return true
	}
	if conNamespace == "au" && typeNamespace == "a" {
		return true
	}
	if conNamespace == "b" && typeNamespace == "a" {
		return true
	}
	return false
}

func LegacyEnableWarningsSimpleTypeNameSkip(conFullName string) bool {
	if conFullName == "healthLoyalty.tmpGetCatalogResultOk" {
		return true
	}
	if conFullName == "lists2.truncateOnReindex" {
		return true
	}
	if conFullName == "gopusher2Vault.wnsStoreCredentials" {
		return true
	}
	if conFullName == "passwordLong.atpTokenSignArgon2Parameters" {
		return true
	}
	if conFullName == "photo.videoLocationFile" {
		return true
	}
	if conFullName == "money.getAccountSummarySuccess" {
		return true
	}
	if conFullName == "healthLoyalty.userTasksGetCollectableCoinsResultOk" {
		return true
	}
	if conFullName == "statshouseApi.queryPointResponse" {
		return true
	}
	if conFullName == "udp.resendRequestExt" {
		return true
	}
	if conFullName == "healthLoyalty.tmpGetShopItemResultOk" {
		return true
	}
	if conFullName == "statshouseApi.chunkResponse" {
		return true
	}
	if conFullName == "lists2.sublistType" {
		return true
	}
	if conFullName == "vkuth.getTokenReponse" {
		return true
	}
	if conFullName == "tree_stats.counterChangePeriodNewValueLong" {
		return true
	}
	if conFullName == "kphp.queueTypesMode" {
		return true
	}
	if conFullName == "messages.oneMarkedMessage" {
		return true
	}
	if conFullName == "money.checkSystemReadyResultSuccess" {
		return true
	}
	if conFullName == "engine.switchBinlogModeSuccess" {
		return true
	}
	if conFullName == "statshouseApi.queryResponse" {
		return true
	}
	if conFullName == "kphp.queueTypesModeV2" {
		return true
	}
	if conFullName == "ch_proxy.table_stats" {
		return true
	}
	if conFullName == "healthLoyalty.tmpGetTransactionsHistoryResultOk" {
		return true
	}
	if conFullName == "engine.metafilesStatData" {
		return true
	}
	if conFullName == "healthLoyalty.tmpBuyShopItemResultOk" {
		return true
	}
	if conFullName == "healthLoyalty.sagaCollectAllRewardsOk" {
		return true
	}
	if conFullName == "healthLoyalty.tmpGetBalanceResultOk" {
		return true
	}
	if conFullName == "healthLoyalty.tmpGetInventoryResultOk" {
		return true
	}
	if conFullName == "healthLoyalty.tmpGetTransactionsResultOk" {
		return true
	}
	if conFullName == "urlBoss.audioNewUrl404Params" {
		return true
	}
	return false
}

// too many of them, impractical
func LegacyEnableWarningsUnionNamePrefixSkip(conName string, typePrefix string, typeSuffix string) bool {
	return true // skip all warnings for now
}

func LegacyEnableWarningsUnionNameExactSkip(conFullName string) bool {
	if conFullName == "engine.queryResult" {
		return true
	}
	if conFullName == "rpcReqResult" {
		return true
	}
	if conFullName == "storage.fileContent" {
		return true
	}
	if conFullName == "tls.combinator" {
		return true
	}
	if conFullName == "tls.combinatorLeft" {
		return true
	}
	if conFullName == "tls.typeExpr" {
		return true
	}
	// below is goldmaster, we keep it to test generated code for this case
	if conFullName == "cd.response" {
		return true
	}
	if conFullName == "a.color" {
		return true
	}
	return false
}

func LegacyGenerateUnusedNatTemplates(conFullName string) bool {
	return conFullName == "rpcInvokeReqExtra" || conFullName == "rpcReqResultExtra"
}

func LegacyAllowTrueBoxed(conFullName string, fieldName string) bool {
	if conFullName == "genericModel.predict" && fieldName == "need_results_float" {
		return true
	}
	if conFullName == "gopusher2.apnsPush" && fieldName == "is_voip" {
		return true
	}
	if conFullName == "storage2.namespacePolicyBlocksCreation" && fieldName == "enable_blocks_creation" {
		return true
	}
	if conFullName == "urlBoss.impParams" && fieldName == "keepAspectRatio" {
		return true
	}
	if conFullName == "urlBoss.impParams" && fieldName == "proxy" {
		return true
	}
	if conFullName == "urlBoss2.impParams" && fieldName == "keepAspectRatio" {
		return true
	}
	if conFullName == "urlBoss2.impParams" && fieldName == "proxy" {
		return true
	}
	if conFullName == "urlBoss2.routingParams" && fieldName == "randomSeed" {
		return true
	}
	// below is goldmaster, we keep it to test generated code for this case
	if conFullName == "useTrue" && fieldName == "b" {
		return true
	}
	if conFullName == "useTrue" && fieldName == "d" {
		return true
	}
	return false
}

func LegacyAllowBoolFieldsmask(conFullName string, fieldName string) bool {
	if conFullName == "adsLalProcessing.processing" && fieldName == "is_deleted" {
		return true
	}
	if conFullName == "adsLalProcessing.startProcessing" && fieldName == "use_apps" {
		return true
	}
	if conFullName == "adsLalProcessing.startProcessing" && fieldName == "use_keywords" {
		return true
	}
	if conFullName == "audiofpgen.genToPMCEx" && fieldName == "allow_short" {
		return true
	}
	if conFullName == "audiofpgen.genToPMCEx" && fieldName == "delete_after_gen" {
		return true
	}
	if conFullName == "blockchainNftCli.getNftDataRequestItem" && fieldName == "img_update" {
		return true
	}
	if conFullName == "ch_proxy.params" && fieldName == "as_post" {
		return true
	}
	if conFullName == "donutLevels.getResponse" && fieldName == "can_add" {
		return true
	}
	if conFullName == "messages.keyboard" && fieldName == "inline" {
		return true
	}
	if conFullName == "messages.keyboard" && fieldName == "one_time" {
		return true
	}
	if conFullName == "messages.setPeerNotificationsSettings" && fieldName == "muted" {
		return true
	}
	if conFullName == "messagesChat.setNotificationsStatus" && fieldName == "is_push_muted" {
		return true
	}
	if conFullName == "messagesChatLong.changeChatMembers" && fieldName == "response_idlong" {
		return true
	}
	if conFullName == "messagesChatLong.setNotificationsStatus" && fieldName == "is_push_muted" {
		return true
	}
	if conFullName == "messagesLong.setFolderNotificationsSettings" && fieldName == "muted" {
		return true
	}
	if conFullName == "messagesLong.setPeerNotificationsSettings" && fieldName == "muted" {
		return true
	}
	if conFullName == "news2.joinedNews" && fieldName == "hidden_by_privacy" {
		return true
	}
	if conFullName == "news2.modifiedNewsEntry" && fieldName == "hidden_by_privacy" {
		return true
	}
	if conFullName == "news2.typeSettings" && fieldName == "deleted" {
		return true
	}
	if conFullName == "news2.typeSettings" && fieldName == "non_std_lengths" {
		return true
	}
	if conFullName == "news2.typeSettings" && fieldName == "normalize_id_lengths" {
		return true
	}
	if conFullName == "online.setFriendOnlineExtraTyped" && fieldName == "invisible_mode" {
		return true
	}
	if conFullName == "sandbox.attachProcessSettings" && fieldName == "store_stats" {
		return true
	}
	if conFullName == "sandbox.createProcessSettings" && fieldName == "store_stats" {
		return true
	}
	if conFullName == "sandbox.storeResultSettings" && fieldName == "store_stderr" {
		return true
	}
	if conFullName == "sandbox.storeResultSettings" && fieldName == "store_stdout" {
		return true
	}
	if conFullName == "searchService.searchMarket" && fieldName == "has_video" {
		return true
	}
	if conFullName == "service.adsTargAdEngineInfo" && fieldName == "user_ad_has_click" {
		return true
	}
	if conFullName == "socket.operateResponse" && fieldName == "closed" {
		return true
	}
	if conFullName == "storage2Impl.engineState" && fieldName == "too_many_uploads" {
		return true
	}
	if conFullName == "targ.ad" && fieldName == "enabled" {
		return true
	}
	if conFullName == "targ.ad" && fieldName == "result_active" {
		return true
	}
	if conFullName == "targ.ad" && fieldName == "schedule_active" {
		return true
	}
	if conFullName == "targ.ad" && fieldName == "suspended" {
		return true
	}
	if conFullName == "targ.adCampaign" && fieldName == "enabled" {
		return true
	}
	if conFullName == "targ.banner" && fieldName == "enabled" {
		return true
	}
	if conFullName == "targ.lal" && fieldName == "is_user_scores_fill_finished" {
		return true
	}
	if conFullName == "targ.searchAdsFilter" && fieldName == "enabled_status" {
		return true
	}
	if conFullName == "targ.searchAdsFilter" && fieldName == "suspended_status" {
		return true
	}
	if conFullName == "targ.user" && fieldName == "has_photo" {
		return true
	}
	if conFullName == "targ.user" && fieldName == "hidden" {
		return true
	}
	if conFullName == "targ.user" && fieldName == "is_online" {
		return true
	}
	if conFullName == "targ.user" && fieldName == "online_invisible" {
		return true
	}
	if conFullName == "targ.user" && fieldName == "pays_money" {
		return true
	}
	if conFullName == "targ.user" && fieldName == "uses_apps" {
		return true
	}
	if conFullName == "targ.userAdApplyViewOptions" && fieldName == "use_for_predictions" {
		return true
	}
	if conFullName == "targLong.user" && fieldName == "has_photo" {
		return true
	}
	if conFullName == "targLong.user" && fieldName == "hidden" {
		return true
	}
	if conFullName == "targLong.user" && fieldName == "is_online" {
		return true
	}
	if conFullName == "targLong.user" && fieldName == "online_invisible" {
		return true
	}
	if conFullName == "targLong.user" && fieldName == "pays_money" {
		return true
	}
	if conFullName == "targLong.user" && fieldName == "uses_apps" {
		return true
	}
	if conFullName == "targUser.adDetails" && fieldName == "is_fake" {
		return true
	}
	if conFullName == "targUser.campaignDetails" && fieldName == "is_fake" {
		return true
	}
	if conFullName == "tasks.queueTypeSettings" && fieldName == "is_blocking" {
		return true
	}
	if conFullName == "tasks.queueTypeSettings" && fieldName == "is_enabled" {
		return true
	}
	if conFullName == "tasks.queueTypeSettings" && fieldName == "is_persistent" {
		return true
	}
	if conFullName == "tasks.queueTypeSettings" && fieldName == "is_staging" {
		return true
	}
	if conFullName == "tasks.queueTypeSettings" && fieldName == "move_to_queue_begin_on_retry" {
		return true
	}
	if conFullName == "wall.searchV2" && fieldName == "reverse" {
		return true
	}
	// below is goldmaster, we keep it to test generated code for this case
	if conFullName == "useTrue" && fieldName == "e" {
		return true
	}
	// legacyExceptions = append(legacyExceptions, legacyPair{conName: conFullName, fieldName: fieldName})
	return false
}
