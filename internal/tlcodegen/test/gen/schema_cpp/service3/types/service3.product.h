#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tl2 { namespace service3 { 
struct Product {
	int32_t type = 0;
	std::vector<int32_t> id;
	std::vector<int32_t> info;
	int32_t date = 0;
	int32_t expiration_date = 0;
	bool removed = false;

	std::string_view tl_name() const { return "service3.product"; }
	uint32_t tl_tag() const { return 0x461f4ce2; }

	bool write_json(std::ostream& s, uint32_t nat_mode)const;

	bool read(::basictl::tl_istream & s, uint32_t nat_mode) noexcept;
	bool write(::basictl::tl_ostream & s, uint32_t nat_mode)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_mode);
	void write_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_mode)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_mode) noexcept;
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_mode)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_mode);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_mode)const;
};

}} // namespace tl2::service3

namespace tl2 { namespace service3 { 
template<uint32_t mode>
struct Productmode {
	int32_t type = 0;
	std::vector<int32_t> id;
	std::vector<int32_t> info;
	int32_t date = 0;
	int32_t expiration_date = 0;
	bool removed = false;

	std::string_view tl_name() const { return "service3.product"; }
	uint32_t tl_tag() const { return 0x461f4ce2; }
};

}} // namespace tl2::service3

