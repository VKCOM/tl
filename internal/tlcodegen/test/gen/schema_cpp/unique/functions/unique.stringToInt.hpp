#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace unique { 
struct StringToInt {
	std::string key;

	std::string_view tl_name() const { return "unique.stringToInt"; }
	uint32_t tl_tag() const { return 0x0f766c35; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, int32_t & result);
	bool write_result(::basictl::tl_ostream & s, int32_t & result);

	friend std::ostream& operator<<(std::ostream& s, const StringToInt& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::unique

