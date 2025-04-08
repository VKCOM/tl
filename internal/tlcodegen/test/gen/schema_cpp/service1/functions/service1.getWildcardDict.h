#pragma once

#include "../../basictl/io_streams.h"
#include "../../__common_namespace/types/dictionary.h"


namespace tl2 { namespace service1 { 
struct GetWildcardDict {
	std::string prefix;

	std::string_view tl_name() const { return "service1.getWildcardDict"; }
	uint32_t tl_tag() const { return 0x72bbc81b; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::Dictionary<std::string> & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::Dictionary<std::string> & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::Dictionary<std::string> & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::Dictionary<std::string> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetWildcardDict& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

