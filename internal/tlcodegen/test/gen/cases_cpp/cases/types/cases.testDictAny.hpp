#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/dictionaryAny.hpp"


namespace tl2 { namespace cases { 
struct TestDictAny {
	::tl2::DictionaryAny<double, int32_t> dict{};

	std::string_view tl_name() const { return "cases.testDictAny"; }
	uint32_t tl_tag() const { return 0xe29b8ae6; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::cases

