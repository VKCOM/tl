#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { 
struct MyString {
	std::string val2;

	std::string_view tl_name() const { return "myString"; }
	uint32_t tl_tag() const { return 0xc8bfa969; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const MyString& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

