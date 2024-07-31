#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../functions/unique.get.hpp"
#include "../../__common_namespace/types/int.hpp"

namespace tl2 { namespace details { 

void UniqueGetReset(::tl2::unique::Get& item);

bool UniqueGetWriteJSON(std::ostream& s, const ::tl2::unique::Get& item);
bool UniqueGetRead(::basictl::tl_istream & s, ::tl2::unique::Get& item);
bool UniqueGetWrite(::basictl::tl_ostream & s, const ::tl2::unique::Get& item);
bool UniqueGetReadBoxed(::basictl::tl_istream & s, ::tl2::unique::Get& item);
bool UniqueGetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::unique::Get& item);

bool UniqueGetReadResult(::basictl::tl_istream & s, ::tl2::unique::Get& item, std::optional<int32_t>& result);
bool UniqueGetWriteResult(::basictl::tl_ostream & s, ::tl2::unique::Get& item, std::optional<int32_t>& result);
		
}} // namespace tl2::details

