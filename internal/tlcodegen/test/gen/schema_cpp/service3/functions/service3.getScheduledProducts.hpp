#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service3.product.hpp"


namespace tl2 { namespace service3 { 
struct GetScheduledProducts {
	int32_t user_id = 0;
	std::vector<int32_t> types;

	std::string_view tl_name() const { return "service3.getScheduledProducts"; }
	uint32_t tl_tag() const { return 0xf53ad7bd; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetScheduledProducts& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

