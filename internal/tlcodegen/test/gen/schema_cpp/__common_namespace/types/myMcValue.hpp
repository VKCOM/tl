#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../service1/types/service1.Value.hpp"


namespace tl2 { 
struct MyMcValue {
	::tl2::service1::Value x;

	std::string_view tl_name() const { return "myMcValue"; }
	uint32_t tl_tag() const { return 0xe2ffd978; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const MyMcValue& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

