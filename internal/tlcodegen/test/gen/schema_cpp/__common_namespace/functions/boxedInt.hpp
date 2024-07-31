#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct BoxedInt {
	int32_t x = 0;

	std::string_view tl_name() const { return "boxedInt"; }
	uint32_t tl_tag() const { return 0x5688ebaf; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, int32_t & result);
	bool write_result(::basictl::tl_ostream & s, int32_t & result);

	friend std::ostream& operator<<(std::ostream& s, const BoxedInt& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

