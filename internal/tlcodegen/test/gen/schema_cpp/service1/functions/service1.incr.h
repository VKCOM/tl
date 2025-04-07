#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service1.Value.h"


namespace tl2 { namespace service1 { 
struct Incr {
	std::string key;
	int64_t value = 0;

	std::string_view tl_name() const { return "service1.incr"; }
	uint32_t tl_tag() const { return 0x0f96b56e; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::service1::Value & result) noexcept;
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::service1::Value & result) noexcept;

	friend std::ostream& operator<<(std::ostream& s, const Incr& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

