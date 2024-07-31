#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service5.params.hpp"
#include "../types/service5.Output.hpp"


namespace tl2 { namespace service5 { 
struct Query {
	std::string query;
	::tl2::service5::Params params{};

	std::string_view tl_name() const { return "service5.query"; }
	uint32_t tl_tag() const { return 0xb3b62513; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::service5::Output & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::service5::Output & result);

	friend std::ostream& operator<<(std::ostream& s, const Query& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service5

