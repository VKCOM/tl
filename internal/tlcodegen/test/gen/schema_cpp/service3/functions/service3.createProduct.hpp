#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service3 { 
struct CreateProduct {
	int32_t user_id = 0;
	int32_t type = 0;
	std::vector<int32_t> id;
	std::vector<int32_t> info;
	int32_t date = 0;
	int32_t expiration_date = 0;

	std::string_view tl_name() const { return "service3.createProduct"; }
	uint32_t tl_tag() const { return 0xb7d92bd9; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);
};

}} // namespace tl2::service3

