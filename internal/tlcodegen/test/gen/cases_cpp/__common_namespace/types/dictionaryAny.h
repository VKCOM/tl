#pragma once

#include "../../basictl/io_streams.h"
#include "dictionaryFieldAny.h"


namespace tl2 { 
template<typename k, typename v>
using DictionaryAny = std::vector<::tl2::DictionaryFieldAny<k, v>>;
} // namespace tl2

