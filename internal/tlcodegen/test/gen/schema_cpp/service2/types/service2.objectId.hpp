#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace service2 { 
struct ObjectId {
	std::vector<int32_t> id;

	std::string_view tl_name() const { return "service2.objectId"; }
	uint32_t tl_tag() const { return 0xaa0af282; }

	bool write_json(std::ostream& s, uint32_t nat_objectIdLength)const;

	bool read(::basictl::tl_istream & s, uint32_t nat_objectIdLength);
	bool write(::basictl::tl_ostream & s, uint32_t nat_objectIdLength)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_objectIdLength);
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_objectIdLength)const;
};

}} // namespace tl2::service2

