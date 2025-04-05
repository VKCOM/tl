#pragma once

#include "../../basictl/io_streams.h"
#include "dictionaryField.h"


namespace tl2 { 
template<typename t>
using Dictionary = std::vector<::tl2::DictionaryField<t>>;
} // namespace tl2

