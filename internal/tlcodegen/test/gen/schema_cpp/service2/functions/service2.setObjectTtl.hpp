#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service2.objectId.hpp"
#include "../../__common/types/true.hpp"


namespace tl2 { namespace service2 { 
struct SetObjectTtl {
	uint32_t objectIdLength = 0;
	::tl2::service2::ObjectId objectId{};
	int32_t ttl = 0;

	std::string_view tl_name() const { return "service2.setObjectTtl"; }
	uint32_t tl_tag() const { return 0x6f98f025; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::True & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::True & result);
};

}} // namespace tl2::service2

