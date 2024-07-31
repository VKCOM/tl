#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../service1/types/service1.Value.hpp"


namespace tl2 { 
struct MyMcValueTuple {
	std::array<::tl2::service1::Value, 3> xs{};

	std::string_view tl_name() const { return "myMcValueTuple"; }
	uint32_t tl_tag() const { return 0x1287d116; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const MyMcValueTuple& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

