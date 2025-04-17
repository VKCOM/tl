#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/right.h"
#include "__common_namespace/types/left.h"


namespace tl2 { 
template<typename X, typename Y>
struct Either {
	std::variant<::tl2::Left<X, Y>, ::tl2::Right<X, Y>> value;

	bool is_left() const { return value.index() == 0; }
	bool is_right() const { return value.index() == 1; }


	std::string_view tl_name() const;
	uint32_t tl_tag() const;
};

} // namespace tl2

