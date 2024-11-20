#pragma once

#include "../../../basics/basictl.h"
#include "../types/service3.productStatsOld.h"


namespace tl2 { namespace service3 { 
struct GetProductStats {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x261f6898; }

	int32_t user_id = 0;
	std::vector<int32_t> types;

	std::string_view tl_name() const { return "service3.getProductStats"; }
	uint32_t tl_tag() const { return 0x261f6898; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::ProductStatsOld>> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<std::vector<::tl2::service3::ProductStatsOld>> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetProductStats& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

