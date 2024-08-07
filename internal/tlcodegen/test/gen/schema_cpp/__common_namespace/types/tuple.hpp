#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
template<typename t>
using Tuple = std::vector<t>;
} // namespace tl2

namespace tl2 { 
template<typename t, uint32_t n>
using Tuplen = std::array<t, n>;
} // namespace tl2

