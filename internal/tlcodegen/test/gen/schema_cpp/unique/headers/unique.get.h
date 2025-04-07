#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../functions/unique.get.h"
#include "../../__common_namespace/types/int.h"

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

