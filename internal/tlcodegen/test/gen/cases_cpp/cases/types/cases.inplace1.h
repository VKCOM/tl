#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.inplace2.h"


namespace tl2 { namespace cases { 
template<typename X>
struct Inplace1 {
	::tl2::cases::Inplace2<X> value{};

	std::string_view tl_name() const { return "cases.inplace1"; }
	uint32_t tl_tag() const { return 0x5533e8e9; }
};

}} // namespace tl2::cases

