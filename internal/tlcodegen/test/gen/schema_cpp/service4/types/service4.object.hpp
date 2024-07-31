#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service4 { 
struct Object {
	int32_t type = 0;
	std::vector<int32_t> joint_id;
	std::vector<int32_t> object_id;

	std::string_view tl_name() const { return "service4.object"; }
	uint32_t tl_tag() const { return 0xa6eeca4f; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const Object& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service4

