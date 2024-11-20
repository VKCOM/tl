#pragma once

#include "../../../basics/basictl.h"
#include "../types/service1.Value.h"


namespace tl2 { namespace service1 { 
struct Get {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x29099b19; }

	std::string key;

	std::string_view tl_name() const { return "service1.get"; }
	uint32_t tl_tag() const { return 0x29099b19; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result);

	friend std::ostream& operator<<(std::ostream& s, const Get& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service1

