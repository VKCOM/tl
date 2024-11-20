#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace service3 { 
struct CreateProduct {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0xb7d92bd9; }

	int32_t user_id = 0;
	int32_t type = 0;
	std::vector<int32_t> id;
	std::vector<int32_t> info;
	int32_t date = 0;
	int32_t expiration_date = 0;

	std::string_view tl_name() const { return "service3.createProduct"; }
	uint32_t tl_tag() const { return 0xb7d92bd9; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);

	friend std::ostream& operator<<(std::ostream& s, const CreateProduct& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

