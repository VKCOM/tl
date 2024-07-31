#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service5.Output.hpp"


namespace tl2 { namespace service5 { 
struct PerformQuery {
	std::string query;

	std::string_view tl_name() const { return "service5.performQuery"; }
	uint32_t tl_tag() const { return 0x019d80a5; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::service5::Output & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::service5::Output & result);

	friend std::ostream& operator<<(std::ostream& s, const PerformQuery& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service5

