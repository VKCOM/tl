#pragma once

#include "../../../basics/basictl.h"
#include "dictionaryField.h"


namespace tl2 { 
template<typename t>
using Dictionary = std::vector<::tl2::DictionaryField<t>>;
} // namespace tl2

