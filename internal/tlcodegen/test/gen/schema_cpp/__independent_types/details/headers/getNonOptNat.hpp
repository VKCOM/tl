#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/functions/getNonOptNat.hpp"
#include "../../../__common/types/int.hpp"

namespace tl2 { namespace details { 

void GetNonOptNatReset(::tl2::GetNonOptNat& item);
bool GetNonOptNatRead(::basictl::tl_istream & s, ::tl2::GetNonOptNat& item);
bool GetNonOptNatWrite(::basictl::tl_ostream & s, const ::tl2::GetNonOptNat& item);
bool GetNonOptNatReadBoxed(::basictl::tl_istream & s, ::tl2::GetNonOptNat& item);
bool GetNonOptNatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetNonOptNat& item);

bool GetNonOptNatReadResult(::basictl::tl_istream & s, ::tl2::GetNonOptNat& item, std::vector<int32_t>& result);
bool GetNonOptNatWriteResult(::basictl::tl_ostream & s, ::tl2::GetNonOptNat& item, std::vector<int32_t>& result);
		
}} // namespace tl2::details

