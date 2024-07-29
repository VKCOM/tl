#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "dictionaryFieldAny.hpp"


namespace tl2 { 
template<typename k, typename v>
using DictionaryAny = std::vector<::tl2::DictionaryFieldAny<k, v>>;
} // namespace tl2

