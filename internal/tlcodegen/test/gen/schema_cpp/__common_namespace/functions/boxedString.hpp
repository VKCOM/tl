#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { 
struct BoxedString {
	// tl magic for function
	static const uint32_t MAGIC = 0x548994db;

	std::string x;

	std::string_view tl_name() const { return "boxedString"; }
	uint32_t tl_tag() const { return 0x548994db; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::string & result);
	bool write_result(::basictl::tl_ostream & s, std::string & result);

	friend std::ostream& operator<<(std::ostream& s, const BoxedString& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

