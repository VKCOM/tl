#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/rpcInvokeReqExtra.h"

namespace tl2 { namespace details { 

void RpcInvokeReqExtraReset(::tl2::RpcInvokeReqExtra& item) noexcept;

bool RpcInvokeReqExtraWriteJSON(std::ostream& s, const ::tl2::RpcInvokeReqExtra& item) noexcept;
bool RpcInvokeReqExtraRead(::basictl::tl_istream & s, ::tl2::RpcInvokeReqExtra& item) noexcept; 
bool RpcInvokeReqExtraWrite(::basictl::tl_ostream & s, const ::tl2::RpcInvokeReqExtra& item) noexcept;
bool RpcInvokeReqExtraReadBoxed(::basictl::tl_istream & s, ::tl2::RpcInvokeReqExtra& item);
bool RpcInvokeReqExtraWriteBoxed(::basictl::tl_ostream & s, const ::tl2::RpcInvokeReqExtra& item);

}} // namespace tl2::details

