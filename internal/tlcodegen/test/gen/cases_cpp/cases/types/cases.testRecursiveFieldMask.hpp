#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/true.hpp"


namespace tl2 { namespace cases { 
struct TestRecursiveFieldMask {
	uint32_t f0 = 0;
	uint32_t f1 = 0;
	uint32_t f2 = 0;
	::tl2::True t1{};
	::tl2::True t2{};
	::tl2::True t3{};

	std::string_view tl_name() const { return "cases.testRecursiveFieldMask"; }
	uint32_t tl_tag() const { return 0xc58cf85e; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

