#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service3 { 
struct DeleteAllProducts {
	int32_t user_id = 0;
	int32_t type = 0;
	int32_t start_date = 0;
	int32_t end_date = 0;

	std::string_view tl_name() const { return "service3.deleteAllProducts"; }
	uint32_t tl_tag() const { return 0x4494acc2; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);
};

}} // namespace tl2::service3

