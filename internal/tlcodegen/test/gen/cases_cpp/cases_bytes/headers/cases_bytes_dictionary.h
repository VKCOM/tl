#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../../__common_namespace/types/dictionary.h"

namespace tl2 { namespace details { 

void DictionaryStringReset(::tl2::Dictionary<std::string>& item);

bool DictionaryStringWriteJSON(std::ostream& s, const ::tl2::Dictionary<std::string>& item);
bool DictionaryStringRead(::basictl::tl_istream & s, ::tl2::Dictionary<std::string>& item);
bool DictionaryStringWrite(::basictl::tl_ostream & s, const ::tl2::Dictionary<std::string>& item);
bool DictionaryStringReadBoxed(::basictl::tl_istream & s, ::tl2::Dictionary<std::string>& item);
bool DictionaryStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Dictionary<std::string>& item);

}} // namespace tl2::details

