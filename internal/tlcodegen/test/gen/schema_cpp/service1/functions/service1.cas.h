#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace service1 { 
struct Cas {
	std::string key;
	int32_t flags = 0;
	int32_t delay = 0;
	std::string casToken;
	std::string newValue;

	std::string_view tl_name() const { return "service1.cas"; }
	uint32_t tl_tag() const { return 0x51851964; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, bool & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result);

	friend std::ostream& operator<<(std::ostream& s, const Cas& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

