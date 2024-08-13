#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service3 { 
struct DeleteProduct {
	// tl magic for function
	static const uint32_t MAGIC = 0x6867e707;

	int32_t user_id = 0;
	int32_t type = 0;
	std::vector<int32_t> id;
	std::vector<int32_t> info;

	std::string_view tl_name() const { return "service3.deleteProduct"; }
	uint32_t tl_tag() const { return 0x6867e707; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result);
	bool write_result(::basictl::tl_ostream & s, bool & result);

	friend std::ostream& operator<<(std::ostream& s, const DeleteProduct& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

