#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service3.limits.h"


namespace tl2 { namespace service3 { 
struct GetLimits {

	std::string_view tl_name() const { return "service3.getLimits"; }
	uint32_t tl_tag() const { return 0xeb399467; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::service3::Limits & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::service3::Limits & result);

	friend std::ostream& operator<<(std::ostream& s, const GetLimits& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

