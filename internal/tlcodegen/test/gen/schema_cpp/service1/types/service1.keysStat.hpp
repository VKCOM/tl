#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/dictionary.hpp"


namespace tl2 { namespace service1 { 
struct KeysStat {
	int32_t start_time = 0;
	::tl2::Dictionary<::tl2::Dictionary<int32_t>> keys_tops{};

	std::string_view tl_name() const { return "service1.keysStat"; }
	uint32_t tl_tag() const { return 0xf0f6bc68; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const KeysStat& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

