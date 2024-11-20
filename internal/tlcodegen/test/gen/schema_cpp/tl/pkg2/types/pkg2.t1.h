#pragma once

#include "../../../basics/basictl.h"
#include "pkg2.foo.h"


namespace tl2 { namespace pkg2 { 
struct T1 {
	::tl2::pkg2::Foo x{};

	std::string_view tl_name() const { return "pkg2.t1"; }
	uint32_t tl_tag() const { return 0x638206ec; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const T1& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::pkg2

