#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../service1/types/service1.Value.hpp"


namespace tl2 { 
struct MyMcValueVector {
	std::vector<::tl2::service1::Value> xs;

	std::string_view tl_name() const { return "myMcValueVector"; }
	uint32_t tl_tag() const { return 0x761d6d58; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

