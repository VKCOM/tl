#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "myString.hpp"
#include "myInt.hpp"


namespace tl2 { 
struct MyValue {
	std::variant<::tl2::MyInt, ::tl2::MyString> value;

	bool is_myInt() const { return value.index() == 0; }
	bool is_myString() const { return value.index() == 1; }


	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool write_json(std::ostream& s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

