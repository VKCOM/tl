#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { 
struct BoxedTuple {
	std::array<int32_t, 3> x{};

	std::string_view tl_name() const { return "boxedTuple"; }
	uint32_t tl_tag() const { return 0x30c9d533; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::array<int32_t, 3> & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, std::array<int32_t, 3> & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, std::array<int32_t, 3> & result) noexcept;
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, std::array<int32_t, 3> & result) noexcept;

	friend std::ostream& operator<<(std::ostream& s, const BoxedTuple& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

