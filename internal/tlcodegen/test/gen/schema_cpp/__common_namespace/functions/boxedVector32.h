#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { 
struct BoxedVector32 {
	std::vector<int32_t> x;

	std::string_view tl_name() const { return "boxedVector32"; }
	uint32_t tl_tag() const { return 0xbbadef07; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result);

	friend std::ostream& operator<<(std::ostream& s, const BoxedVector32& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

