#pragma once

#include "../../basictl/io_streams.h"
#include "../../__common_namespace/types/pair.h"


namespace tl2 { namespace cases { 
template<typename X>
struct Inplace3 {
	::tl2::Pair<std::vector<X>, std::vector<X>> value{};

	std::string_view tl_name() const { return "cases.inplace3"; }
	uint32_t tl_tag() const { return 0x4ffb95cb; }
};

}} // namespace tl2::cases

