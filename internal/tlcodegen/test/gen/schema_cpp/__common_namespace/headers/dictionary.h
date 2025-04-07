#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/dictionary.h"

namespace tl2 { namespace details { 

void DictionaryIntReset(::tl2::Dictionary<int32_t>& item);

bool DictionaryIntWriteJSON(std::ostream& s, const ::tl2::Dictionary<int32_t>& item);
bool DictionaryIntRead(::basictl::tl_istream & s, ::tl2::Dictionary<int32_t>& item);
bool DictionaryIntWrite(::basictl::tl_ostream & s, const ::tl2::Dictionary<int32_t>& item);
bool DictionaryIntReadBoxed(::basictl::tl_istream & s, ::tl2::Dictionary<int32_t>& item);
bool DictionaryIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Dictionary<int32_t>& item);

}} // namespace tl2::details

