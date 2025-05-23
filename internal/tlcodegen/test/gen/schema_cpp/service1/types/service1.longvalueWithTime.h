// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tl2 { namespace service1 { 
struct LongvalueWithTime {
	int64_t value = 0;
	int32_t flags = 0;
	int32_t modificationTime = 0;

	std::string_view tl_name() const { return "service1.longvalueWithTime"; }
	uint32_t tl_tag() const { return 0xa04606ec; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const LongvalueWithTime& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

