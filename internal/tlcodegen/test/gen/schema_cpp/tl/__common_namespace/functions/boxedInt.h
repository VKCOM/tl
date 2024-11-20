#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { 
struct BoxedInt {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x5688ebaf; }

	int32_t x = 0;

	std::string_view tl_name() const { return "boxedInt"; }
	uint32_t tl_tag() const { return 0x5688ebaf; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, int32_t & result);
	bool write_result(::basictl::tl_ostream & s, int32_t & result);

	friend std::ostream& operator<<(std::ostream& s, const BoxedInt& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

