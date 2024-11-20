#pragma once

#include "../../../basics/basictl.h"
#include "../types/myDouble.h"


namespace tl2 { 
struct GetMyDouble {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0xb660ad10; }

	::tl2::MyDouble x{};

	std::string_view tl_name() const { return "getMyDouble"; }
	uint32_t tl_tag() const { return 0xb660ad10; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::MyDouble & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::MyDouble & result);

	friend std::ostream& operator<<(std::ostream& s, const GetMyDouble& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

