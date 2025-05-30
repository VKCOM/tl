// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/UnionArgsXXX.h"


namespace tl2 { 
struct UnionArgsUse {
	uint32_t k = 0;
	uint32_t n = 0;
	::tl2::UnionArgsXXX<int32_t> a;
	::tl2::UnionArgsXXX<int64_t> b;

	std::string_view tl_name() const { return "unionArgsUse"; }
	uint32_t tl_tag() const { return 0x742161d2; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const UnionArgsUse& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

