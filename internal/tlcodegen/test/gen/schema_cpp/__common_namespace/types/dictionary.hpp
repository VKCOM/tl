#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "dictionaryField.hpp"


namespace tl2 { 
template<typename t>
using Dictionary = std::vector<::tl2::DictionaryField<t>>;
} // namespace tl2

