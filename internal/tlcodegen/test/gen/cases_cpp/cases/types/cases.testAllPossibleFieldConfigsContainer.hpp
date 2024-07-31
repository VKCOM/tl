#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "cases.testAllPossibleFieldConfigs.hpp"


namespace tl2 { namespace cases { 
struct TestAllPossibleFieldConfigsContainer {
	uint32_t outer = 0;
	::tl2::cases::TestAllPossibleFieldConfigs value{};

	std::string_view tl_name() const { return "cases.testAllPossibleFieldConfigsContainer"; }
	uint32_t tl_tag() const { return 0xe3fae936; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestAllPossibleFieldConfigsContainer& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

