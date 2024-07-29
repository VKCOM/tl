#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "service5.stringOutput.hpp"
#include "service5.emptyOutput.hpp"


namespace tl2 { namespace service5 { 
struct Output {
	std::variant<::tl2::service5::EmptyOutput, ::tl2::service5::StringOutput> value;

	bool is_empty() const { return value.index() == 0; }
	bool is_string() const { return value.index() == 1; }

	void set_empty() { value.emplace<0>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::service5

