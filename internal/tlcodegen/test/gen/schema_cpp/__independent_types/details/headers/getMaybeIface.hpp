#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/functions/getMaybeIface.hpp"
#include "../../../service1/types/service1.Value.hpp"

namespace tl2 { namespace details { 

void GetMaybeIfaceReset(::tl2::GetMaybeIface& item);
bool GetMaybeIfaceRead(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item);
bool GetMaybeIfaceWrite(::basictl::tl_ostream & s, const ::tl2::GetMaybeIface& item);
bool GetMaybeIfaceReadBoxed(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item);
bool GetMaybeIfaceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMaybeIface& item);

bool GetMaybeIfaceReadResult(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item, std::optional<::tl2::service1::Value>& result);
bool GetMaybeIfaceWriteResult(::basictl::tl_ostream & s, ::tl2::GetMaybeIface& item, std::optional<::tl2::service1::Value>& result);
		
}} // namespace tl2::details

