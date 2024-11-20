#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { 
struct BoxedTuple {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x30c9d533; }

	std::array<int32_t, 3> x{};

	std::string_view tl_name() const { return "boxedTuple"; }
	uint32_t tl_tag() const { return 0x30c9d533; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::array<int32_t, 3> & result);
	bool write_result(::basictl::tl_ostream & s, std::array<int32_t, 3> & result);

	friend std::ostream& operator<<(std::ostream& s, const BoxedTuple& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

