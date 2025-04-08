#pragma once

#include "../../basictl/io_streams.h"
#include "../types/myBoxedTupleSlice.h"


namespace tl2 { 
struct BoxedTupleSlice2 {
	::tl2::MyBoxedTupleSlice x{};

	std::string_view tl_name() const { return "boxedTupleSlice2"; }
	uint32_t tl_tag() const { return 0x1cdf4705; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::MyBoxedTupleSlice & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::MyBoxedTupleSlice & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::MyBoxedTupleSlice & result);

	friend std::ostream& operator<<(std::ostream& s, const BoxedTupleSlice2& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

