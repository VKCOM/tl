#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct BoxedTupleSlice3 {
	uint32_t n = 0;
	std::vector<int32_t> x;

	std::string_view tl_name() const { return "boxedTupleSlice3"; }
	uint32_t tl_tag() const { return 0xa19b8106; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result);

	friend std::ostream& operator<<(std::ostream& s, const BoxedTupleSlice3& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

