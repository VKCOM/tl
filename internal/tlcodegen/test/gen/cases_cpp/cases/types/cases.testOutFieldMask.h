#pragma once

#include "../../basictl/io_streams.h"
#include "../../__common_namespace/types/true.h"


namespace tl2 { namespace cases { 
struct TestOutFieldMask {
	uint32_t f1 = 0;
	::tl2::True f2{};
	std::vector<int32_t> f3;

	std::string_view tl_name() const { return "cases.testOutFieldMask"; }
	uint32_t tl_tag() const { return 0xbd6b4b3c; }

	bool write_json(std::ostream& s, uint32_t nat_f)const;

	bool read(::basictl::tl_istream & s, uint32_t nat_f) noexcept;
	bool write(::basictl::tl_ostream & s, uint32_t nat_f)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_f);
	void write_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_f)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_f) noexcept;
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_f)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_f);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_f)const;
};

}} // namespace tl2::cases

