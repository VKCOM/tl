#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myDouble.h"


namespace tl2 { 
struct GetMyDouble {
	::tl2::MyDouble x{};

	std::string_view tl_name() const { return "getMyDouble"; }
	uint32_t tl_tag() const { return 0xb660ad10; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::MyDouble & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::MyDouble & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::MyDouble & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::MyDouble & result);

	friend std::ostream& operator<<(std::ostream& s, const GetMyDouble& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

