// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

// All compatibility code that will be removed after some fixes to TL itself

func EnableWarningsUnionNamespaceSkipLegacy(conNamespace string, typeNamespace string) bool {
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

func EnableWarningsSimpleTypeNameSkipLegacy(conFullName string) bool {
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

func EnableWarningsUnionNamePrefixSkipLegacy(conName string, typePrefix string, typeSuffix string) bool {
	return true // skip all warnings for now
}

func EnableWarningsUnionNameExactSkipLegacy(conFullName string) bool {
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

func GenerateUnusedNatTemplates(conFullName string) bool {
	return conFullName == "rpcInvokeReqExtra" || conFullName == "rpcReqResultExtra"
}
