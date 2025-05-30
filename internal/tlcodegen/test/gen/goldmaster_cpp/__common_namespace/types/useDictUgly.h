// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/pair.h"
#include "__common_namespace/types/multiPoint.h"
#include "__common_namespace/types/dictionaryElemUgly.h"
#include "__common_namespace/types/dictionaryElem.h"
#include "a/types/a.Color.h"


namespace tl2 { 
struct UseDictUgly {
	uint32_t n = 0;
	std::vector<::tl2::DictionaryElemUgly<int32_t, std::string>> a;
	std::map<uint32_t, std::string> b;
	std::vector<::tl2::DictionaryElem<::tl2::Pair<int32_t, int32_t>, int32_t>> c;
	std::vector<::tl2::DictionaryElem<std::vector<std::string>, int32_t>> d;
	std::vector<::tl2::DictionaryElem<::tl2::Pair<bool, ::tl2::a::Color>, int32_t>> e;
	std::vector<::tl2::DictionaryElem<::tl2::Pair<float, double>, int32_t>> f;
	std::vector<::tl2::DictionaryElem<::tl2::Pair<int32_t, ::tl2::Pair<::tl2::MultiPoint, std::string>>, int32_t>> g;
	std::map<int32_t, ::tl2::Pair<int32_t, int32_t>> x;
	std::map<int64_t, ::tl2::Pair<int32_t, int32_t>> y;
	std::map<std::string, ::tl2::Pair<int32_t, int32_t>> z;

	std::string_view tl_name() const { return "useDictUgly"; }
	uint32_t tl_tag() const { return 0xfb9ce817; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const UseDictUgly& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

