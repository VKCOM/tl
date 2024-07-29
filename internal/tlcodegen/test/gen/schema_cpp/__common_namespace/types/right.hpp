#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
template<typename X, typename Y>
struct Right {
	Y value{};

	std::string_view tl_name() const { return "right"; }
	uint32_t tl_tag() const { return 0xdf3ecb3b; }
};

} // namespace tl2

