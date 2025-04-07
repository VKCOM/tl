#pragma once

#include "../../basictl/io_streams.h"
#include "benchmarks.vruposition.h"


namespace tl2 { namespace benchmarks { 
struct VrutoytopLevelUnionBig {
	std::vector<::tl2::benchmarks::Vruposition> next_positions;

	std::string_view tl_name() const { return "benchmarks.vrutoytopLevelUnionBig"; }
	uint32_t tl_tag() const { return 0xef556bee; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const VrutoytopLevelUnionBig& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::benchmarks

