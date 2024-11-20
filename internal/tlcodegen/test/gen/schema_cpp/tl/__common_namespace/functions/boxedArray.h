#pragma once

#include "../../../basics/basictl.h"
#include "../types/myBoxedArray.h"


namespace tl2 { 
struct BoxedArray {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x95dcc8b7; }

	::tl2::MyBoxedArray x{};

	std::string_view tl_name() const { return "boxedArray"; }
	uint32_t tl_tag() const { return 0x95dcc8b7; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::MyBoxedArray & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::MyBoxedArray & result);

	friend std::ostream& operator<<(std::ostream& s, const BoxedArray& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

