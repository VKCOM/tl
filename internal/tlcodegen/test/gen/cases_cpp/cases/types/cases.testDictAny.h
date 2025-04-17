#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/dictionaryAny.h"


namespace tl2 { namespace cases { 
struct TestDictAny {
	::tl2::DictionaryAny<double, int32_t> dict{};

	std::string_view tl_name() const { return "cases.testDictAny"; }
	uint32_t tl_tag() const { return 0xe29b8ae6; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TestDictAny& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::cases

