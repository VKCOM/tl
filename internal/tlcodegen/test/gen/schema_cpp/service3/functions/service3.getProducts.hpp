#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service3.product.hpp"


namespace tl2 { namespace service3 { 
struct GetProducts {
	int32_t user_id = 0;
	uint32_t mode = 0;
	std::vector<int32_t> types;
	int32_t start_date = 0;
	int32_t end_date = 0;
	int32_t offset = 0;
	int32_t limit = 0;
	std::vector<int32_t> allowed_info0;

	std::string_view tl_name() const { return "service3.getProducts"; }
	uint32_t tl_tag() const { return 0xeb306233; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Product>> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<std::vector<::tl2::service3::Product>> & result);
};

}} // namespace tl2::service3

