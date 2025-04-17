#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service2/types/service2.objectId.h"
#include "__common_namespace/types/true.h"


namespace tl2 { namespace service2 { 
struct SetObjectTtl {
	uint32_t objectIdLength = 0;
	::tl2::service2::ObjectId objectId{};
	int32_t ttl = 0;

	std::string_view tl_name() const { return "service2.setObjectTtl"; }
	uint32_t tl_tag() const { return 0x6f98f025; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::True & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::True & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::True & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::True & result);

	friend std::ostream& operator<<(std::ostream& s, const SetObjectTtl& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service2

