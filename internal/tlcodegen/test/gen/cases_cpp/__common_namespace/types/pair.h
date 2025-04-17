#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tl2 { 
template<typename X, typename Y>
struct Pair {
	X x{};
	Y y{};

	std::string_view tl_name() const { return "pair"; }
	uint32_t tl_tag() const { return 0xf01604df; }
};

} // namespace tl2

