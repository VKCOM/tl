#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace service3 { 
struct GetLastVisitTimestamp {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0x9a4c788d; }

	int32_t user_id = 0;

	std::string_view tl_name() const { return "service3.getLastVisitTimestamp"; }
	uint32_t tl_tag() const { return 0x9a4c788d; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetLastVisitTimestamp& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

