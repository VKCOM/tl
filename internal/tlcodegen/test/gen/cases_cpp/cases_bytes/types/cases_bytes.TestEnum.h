#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases_bytes/types/cases_bytes.TestEnumItems.h"


namespace tl2 { namespace cases_bytes { 
struct TestEnum {
	std::variant<::tl2::cases_bytes::TestEnum1, ::tl2::cases_bytes::TestEnum2, ::tl2::cases_bytes::TestEnum3> value;

	bool is_1() const { return value.index() == 0; }
	bool is_2() const { return value.index() == 1; }
	bool is_3() const { return value.index() == 2; }

	void set_1() { value.emplace<0>(); }
	void set_2() { value.emplace<1>(); }
	void set_3() { value.emplace<2>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool write_json(std::ostream& s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;
};

}} // namespace tl2::cases_bytes

