#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service3.limits.hpp"
#include "../../__common_namespace/types/boolStat.hpp"


namespace tl2 { namespace service3 { 
struct SetLimits {
	// tl magic for function
	static const uint32_t MAGIC = 0x3ad5c19c;

	::tl2::service3::Limits limits{};

	std::string_view tl_name() const { return "service3.setLimits"; }
	uint32_t tl_tag() const { return 0x3ad5c19c; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::BoolStat & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::BoolStat & result);

	friend std::ostream& operator<<(std::ostream& s, const SetLimits& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

