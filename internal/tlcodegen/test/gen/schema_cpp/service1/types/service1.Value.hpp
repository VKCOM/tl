#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "service1.strvalueWithTime.hpp"
#include "service1.strvalue.hpp"
#include "service1.not_found.hpp"
#include "service1.longvalueWithTime.hpp"
#include "service1.longvalue.hpp"


namespace tl2 { namespace service1 { 
struct Value {
	std::variant<::tl2::service1::Not_found, ::tl2::service1::Strvalue, ::tl2::service1::Longvalue, ::tl2::service1::StrvalueWithTime, ::tl2::service1::LongvalueWithTime> value;

	bool is_not_found() const { return value.index() == 0; }
	bool is_strvalue() const { return value.index() == 1; }
	bool is_longvalue() const { return value.index() == 2; }
	bool is_strvalueWithTime() const { return value.index() == 3; }
	bool is_longvalueWithTime() const { return value.index() == 4; }

	void set_not_found() { value.emplace<0>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool write_json(std::ostream& s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::service1

