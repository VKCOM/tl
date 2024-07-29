#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "right.hpp"
#include "left.hpp"


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

