#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service3.limits.h"
#include "../../__common_namespace/types/boolStat.h"


namespace tl2 { namespace service3 { 
struct SetLimits {
	::tl2::service3::Limits limits{};

	std::string_view tl_name() const { return "service3.setLimits"; }
	uint32_t tl_tag() const { return 0x3ad5c19c; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::BoolStat & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::BoolStat & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::BoolStat & result) noexcept;
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::BoolStat & result) noexcept;

	friend std::ostream& operator<<(std::ostream& s, const SetLimits& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

