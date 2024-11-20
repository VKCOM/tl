#pragma once

#include "../../../basics/basictl.h"
#include "dictionaryFieldAny.h"


namespace tl2 { 
template<typename k, typename v>
using DictionaryAny = std::vector<::tl2::DictionaryFieldAny<k, v>>;
} // namespace tl2

