#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/pair.h"
#include "cases/types/cases.inplace1.h"


namespace tl2 { namespace cases { 
struct TestInplaceStructArgs2 {
	uint32_t a1 = 0;
	uint32_t a2 = 0;
	uint32_t a3 = 0;
	::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>> arg{};

	std::string_view tl_name() const { return "cases.testInplaceStructArgs2"; }
	uint32_t tl_tag() const { return 0xaa9f2480; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestInplaceStructArgs2& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

