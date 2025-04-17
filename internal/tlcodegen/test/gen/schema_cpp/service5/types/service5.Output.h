#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service5/types/service5.stringOutput.h"
#include "service5/types/service5.emptyOutput.h"


namespace tl2 { namespace service5 { 
struct Output {
	std::variant<::tl2::service5::EmptyOutput, ::tl2::service5::StringOutput> value;

	bool is_empty() const { return value.index() == 0; }
	bool is_string() const { return value.index() == 1; }

	void set_empty() { value.emplace<0>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool write_json(std::ostream& s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;
};

}} // namespace tl2::service5

