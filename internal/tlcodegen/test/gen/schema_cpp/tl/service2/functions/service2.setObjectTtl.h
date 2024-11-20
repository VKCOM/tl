#pragma once

#include "../../../basics/basictl.h"
#include "../types/service2.objectId.h"
#include "../../__common_namespace/types/true.h"


namespace tl2 { namespace service2 { 
struct SetObjectTtl {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x6f98f025; }

	uint32_t objectIdLength = 0;
	::tl2::service2::ObjectId objectId{};
	int32_t ttl = 0;

	std::string_view tl_name() const { return "service2.setObjectTtl"; }
	uint32_t tl_tag() const { return 0x6f98f025; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::True & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::True & result);

	friend std::ostream& operator<<(std::ostream& s, const SetObjectTtl& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service2

