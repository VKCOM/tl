#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases_bytes { 
struct TestEnum1 {

	std::string_view tl_name() const { return "cases_bytes.testEnum1"; }
	uint32_t tl_tag() const { return 0x58aad3f5; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases_bytes

namespace tl2 { namespace cases_bytes { 
struct TestEnum2 {

	std::string_view tl_name() const { return "cases_bytes.testEnum2"; }
	uint32_t tl_tag() const { return 0x00b47add; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases_bytes

namespace tl2 { namespace cases_bytes { 
struct TestEnum3 {

	std::string_view tl_name() const { return "cases_bytes.testEnum3"; }
	uint32_t tl_tag() const { return 0x81911ffa; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases_bytes

