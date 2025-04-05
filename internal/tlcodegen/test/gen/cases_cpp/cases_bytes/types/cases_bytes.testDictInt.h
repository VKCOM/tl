#pragma once

#include "../../basictl/io_streams.h"
#include "../../__common_namespace/types/dictionaryFieldAny.h"


namespace tl2 { namespace cases_bytes { 
struct TestDictInt {
	std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>> dict;

	std::string_view tl_name() const { return "cases_bytes.testDictInt"; }
	uint32_t tl_tag() const { return 0x453ace07; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestDictInt& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases_bytes

