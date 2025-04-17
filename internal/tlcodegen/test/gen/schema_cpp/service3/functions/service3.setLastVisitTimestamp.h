#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace service3 { 
struct SetLastVisitTimestamp {
	int32_t user_id = 0;
	int32_t timestamp = 0;

	std::string_view tl_name() const { return "service3.setLastVisitTimestamp"; }
	uint32_t tl_tag() const { return 0x7909b020; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, bool & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, bool & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) noexcept;
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) noexcept;

	friend std::ostream& operator<<(std::ostream& s, const SetLastVisitTimestamp& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

