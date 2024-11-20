#pragma once

#include "../../../basics/basictl.h"
#include "../types/rpcInvokeReqExtra.h"

namespace tl2 { namespace details { 

void RpcInvokeReqExtraReset(::tl2::RpcInvokeReqExtra& item);

bool RpcInvokeReqExtraWriteJSON(std::ostream& s, const ::tl2::RpcInvokeReqExtra& item);
bool RpcInvokeReqExtraRead(::basictl::tl_istream & s, ::tl2::RpcInvokeReqExtra& item);
bool RpcInvokeReqExtraWrite(::basictl::tl_ostream & s, const ::tl2::RpcInvokeReqExtra& item);
bool RpcInvokeReqExtraReadBoxed(::basictl::tl_istream & s, ::tl2::RpcInvokeReqExtra& item);
bool RpcInvokeReqExtraWriteBoxed(::basictl::tl_ostream & s, const ::tl2::RpcInvokeReqExtra& item);

}} // namespace tl2::details

