#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/dictionary.hpp"


namespace tl2 { namespace cases_bytes { 
struct TestDictStringString {
	::tl2::Dictionary<std::string> dict{};

	std::string_view tl_name() const { return "cases_bytes.testDictStringString"; }
	uint32_t tl_tag() const { return 0xad69c772; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestDictStringString& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases_bytes

