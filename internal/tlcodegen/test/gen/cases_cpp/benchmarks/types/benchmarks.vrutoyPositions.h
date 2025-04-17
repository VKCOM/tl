#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "benchmarks/types/benchmarks.vruposition.h"


namespace tl2 { namespace benchmarks { 
struct VrutoyPositions {
	std::vector<::tl2::benchmarks::Vruposition> next_positions;

	std::string_view tl_name() const { return "benchmarks.vrutoyPositions"; }
	uint32_t tl_tag() const { return 0xb6003de0; }

	bool write_json(std::ostream& s, uint32_t nat_n)const;

	bool read(::basictl::tl_istream & s, uint32_t nat_n) noexcept;
	bool write(::basictl::tl_ostream & s, uint32_t nat_n)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_n);
	void write_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_n)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_n) noexcept;
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_n)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_n);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_n)const;
};

}} // namespace tl2::benchmarks

