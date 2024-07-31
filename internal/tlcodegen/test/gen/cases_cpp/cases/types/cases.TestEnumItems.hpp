#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace cases { 
struct TestEnum1 {

	std::string_view tl_name() const { return "cases.testEnum1"; }
	uint32_t tl_tag() const { return 0x6c6c55ac; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

namespace tl2 { namespace cases { 
struct TestEnum2 {

	std::string_view tl_name() const { return "cases.testEnum2"; }
	uint32_t tl_tag() const { return 0x86ea88ce; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

namespace tl2 { namespace cases { 
struct TestEnum3 {

	std::string_view tl_name() const { return "cases.testEnum3"; }
	uint32_t tl_tag() const { return 0x69b83e2f; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

