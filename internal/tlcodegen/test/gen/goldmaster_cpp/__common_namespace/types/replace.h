// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/replace9.h"
#include "__common_namespace/types/replace8.h"
#include "__common_namespace/types/replace7.h"
#include "__common_namespace/types/replace6.h"
#include "__common_namespace/types/replace5.h"
#include "__common_namespace/types/replace4.h"
#include "__common_namespace/types/replace3.h"
#include "__common_namespace/types/replace2.h"
#include "__common_namespace/types/replace15.h"
#include "__common_namespace/types/replace14.h"
#include "__common_namespace/types/replace13.h"
#include "__common_namespace/types/replace1.h"
#include "__common_namespace/types/replace12.h"
#include "__common_namespace/types/replace11.h"
#include "__common_namespace/types/replace10.h"


namespace tl2 { 
struct Replace {
	uint32_t n = 0;
	::tl2::Replace1 a{};
	::tl2::Replace1n<3> a1{};
	::tl2::Replace2 b{};
	::tl2::Replace3 c{};
	::tl2::Replace4 d{};
	::tl2::Replace4n<3> d1{};
	::tl2::Replace5 e{};
	::tl2::Replace6 g{};
	::tl2::Replace7 h{};
	::tl2::Replace8 i{};
	::tl2::Replace9 j{};
	::tl2::Replace10 k{};
	::tl2::Replace11<int64_t> l{};
	::tl2::Replace12 m{};
	::tl2::Replace13<int64_t> o{};
	::tl2::Replace14<int64_t> p{};
	::tl2::Replace15 q{};

	std::string_view tl_name() const { return "replace"; }
	uint32_t tl_tag() const { return 0x323db63e; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const Replace& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

