#pragma once

#include "../../basictl/io_streams.h"
#include "../types/myBoxedArray.h"


namespace tl2 { 
struct BoxedArray {
	::tl2::MyBoxedArray x{};

	std::string_view tl_name() const { return "boxedArray"; }
	uint32_t tl_tag() const { return 0x95dcc8b7; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::MyBoxedArray & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::MyBoxedArray & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::MyBoxedArray & result) noexcept;
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::MyBoxedArray & result) noexcept;

	friend std::ostream& operator<<(std::ostream& s, const BoxedArray& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

