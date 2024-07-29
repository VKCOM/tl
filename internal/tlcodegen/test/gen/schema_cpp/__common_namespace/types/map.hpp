#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
template<typename X, typename Y>
struct Map {
	X key{};
	Y value{};

	std::string_view tl_name() const { return "map"; }
	uint32_t tl_tag() const { return 0x79c473a4; }
};

} // namespace tl2

