#pragma once

#include "../../basictl/io_streams.h"
#include "cases.inplace1.h"


namespace tl2 { namespace cases { 
struct TestInplaceStructArgs {
	uint32_t a1 = 0;
	uint32_t a2 = 0;
	uint32_t a3 = 0;
	::tl2::cases::Inplace1<int32_t> arg{};

	std::string_view tl_name() const { return "cases.testInplaceStructArgs"; }
	uint32_t tl_tag() const { return 0xa9e4441e; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestInplaceStructArgs& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

