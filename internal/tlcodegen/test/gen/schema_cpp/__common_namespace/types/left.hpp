#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
template<typename X, typename Y>
struct Left {
	X value{};

	std::string_view tl_name() const { return "left"; }
	uint32_t tl_tag() const { return 0x0a29cd5d; }
};

} // namespace tl2

