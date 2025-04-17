#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/dictionaryFieldAny.h"


namespace tl2 { 
template<typename k, typename v>
using DictionaryAny = std::vector<::tl2::DictionaryFieldAny<k, v>>;
} // namespace tl2

