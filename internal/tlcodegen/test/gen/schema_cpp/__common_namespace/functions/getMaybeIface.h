#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service1/types/service1.Value.h"


namespace tl2 { 
struct GetMaybeIface {
	::tl2::service1::Value x;

	std::string_view tl_name() const { return "getMaybeIface"; }
	uint32_t tl_tag() const { return 0x6b055ae4; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<::tl2::service1::Value> & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, std::optional<::tl2::service1::Value> & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<::tl2::service1::Value> & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<::tl2::service1::Value> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetMaybeIface& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

