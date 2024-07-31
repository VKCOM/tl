#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "dictionary.hpp"


namespace tl2 { 
struct MyTwoDicts {
	::tl2::Dictionary<int32_t> a{};
	::tl2::Dictionary<int32_t> b{};

	std::string_view tl_name() const { return "myTwoDicts"; }
	uint32_t tl_tag() const { return 0xa859581d; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const MyTwoDicts& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

