#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/integer.h"


namespace tl2 { 
struct BenchObject {
	std::vector<int32_t> xs;
	std::vector<::tl2::Integer> ys;

	std::string_view tl_name() const { return "benchObject"; }
	uint32_t tl_tag() const { return 0xb697e865; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const BenchObject& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

