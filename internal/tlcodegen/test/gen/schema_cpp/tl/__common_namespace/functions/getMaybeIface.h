#pragma once

#include "../../../basics/basictl.h"
#include "../../service1/types/service1.Value.h"


namespace tl2 { 
struct GetMaybeIface {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x6b055ae4; }

	::tl2::service1::Value x;

	std::string_view tl_name() const { return "getMaybeIface"; }
	uint32_t tl_tag() const { return 0x6b055ae4; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<::tl2::service1::Value> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<::tl2::service1::Value> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetMaybeIface& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

