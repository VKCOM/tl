#pragma once

#include "../../basictl/io_streams.h"
#include "../../__common_namespace/types/true.h"


namespace tl2 { namespace cases { 
struct TestAllPossibleFieldConfigs {
	uint32_t local = 0;
	int32_t f00 = 0;
	::tl2::True f01{};
	std::vector<int32_t> f02;
	std::vector<int32_t> f03;
	int32_t f10 = 0;
	::tl2::True f11{};
	std::vector<int32_t> f12;
	std::vector<int32_t> f13;
	int32_t f20 = 0;
	::tl2::True f21{};
	std::vector<int32_t> f22;
	std::vector<int32_t> f23;

	std::string_view tl_name() const { return "cases.testAllPossibleFieldConfigs"; }
	uint32_t tl_tag() const { return 0xfb6836d3; }

	bool write_json(std::ostream& s, uint32_t nat_outer)const;

	bool read(::basictl::tl_istream & s, uint32_t nat_outer) noexcept;
	bool write(::basictl::tl_ostream & s, uint32_t nat_outer)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_outer);
	void write_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_outer)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_outer) noexcept;
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_outer)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_outer);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_outer)const;
};

}} // namespace tl2::cases

