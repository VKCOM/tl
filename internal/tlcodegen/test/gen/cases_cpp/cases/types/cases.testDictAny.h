#pragma once

#include "../../basictl/io_streams.h"
#include "../../__common_namespace/types/dictionaryAny.h"


namespace tl2 { namespace cases { 
struct TestDictAny {
	::tl2::DictionaryAny<double, int32_t> dict{};

	std::string_view tl_name() const { return "cases.testDictAny"; }
	uint32_t tl_tag() const { return 0xe29b8ae6; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestDictAny& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

