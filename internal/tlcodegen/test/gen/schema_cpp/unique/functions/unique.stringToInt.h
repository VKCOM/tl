#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace unique { 
struct StringToInt {
	std::string key;

	std::string_view tl_name() const { return "unique.stringToInt"; }
	uint32_t tl_tag() const { return 0x0f766c35; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, int32_t & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, int32_t & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, int32_t & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, int32_t & result);

	friend std::ostream& operator<<(std::ostream& s, const StringToInt& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::unique

