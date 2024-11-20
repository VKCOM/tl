#pragma once

#include "../../../basics/basictl.h"
#include "../functions/getMaybeIface.h"
#include "../../service1/types/service1.Value.h"

namespace tl2 { namespace details { 

void GetMaybeIfaceReset(::tl2::GetMaybeIface& item);

bool GetMaybeIfaceWriteJSON(std::ostream& s, const ::tl2::GetMaybeIface& item);
bool GetMaybeIfaceRead(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item);
bool GetMaybeIfaceWrite(::basictl::tl_ostream & s, const ::tl2::GetMaybeIface& item);
bool GetMaybeIfaceReadBoxed(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item);
bool GetMaybeIfaceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMaybeIface& item);

bool GetMaybeIfaceReadResult(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item, std::optional<::tl2::service1::Value>& result);
bool GetMaybeIfaceWriteResult(::basictl::tl_ostream & s, ::tl2::GetMaybeIface& item, std::optional<::tl2::service1::Value>& result);
		
}} // namespace tl2::details

