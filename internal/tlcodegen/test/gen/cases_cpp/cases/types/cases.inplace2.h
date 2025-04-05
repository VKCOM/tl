#pragma once

#include "../../basictl/io_streams.h"
#include "cases.inplace3.h"


namespace tl2 { namespace cases { 
template<typename X>
struct Inplace2 {
	::tl2::cases::Inplace3<std::array<X, 2>> value{};

	std::string_view tl_name() const { return "cases.inplace2"; }
	uint32_t tl_tag() const { return 0x869fcff5; }
};

}} // namespace tl2::cases

