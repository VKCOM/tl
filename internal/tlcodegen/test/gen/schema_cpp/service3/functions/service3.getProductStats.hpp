#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service3.productStatsOld.hpp"


namespace tl2 { namespace service3 { 
struct GetProductStats {
	int32_t user_id = 0;
	std::vector<int32_t> types;

	std::string_view tl_name() const { return "service3.getProductStats"; }
	uint32_t tl_tag() const { return 0x261f6898; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::ProductStatsOld>> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<std::vector<::tl2::service3::ProductStatsOld>> & result);
};

}} // namespace tl2::service3

