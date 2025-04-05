#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace tasks { 
struct FullFilledCron {
	uint32_t fields_mask = 0;
	int32_t a0 = 0;
	int32_t a1 = 0;
	int32_t a2 = 0;
	int32_t a3 = 0;
	int32_t a4 = 0;
	int32_t a5 = 0;
	int32_t a6 = 0;
	int32_t a7 = 0;
	int32_t a8 = 0;
	int32_t a9 = 0;
	int32_t a10 = 0;
	int32_t a11 = 0;
	int32_t a12 = 0;
	int32_t a13 = 0;
	int32_t a14 = 0;
	int32_t a15 = 0;
	int32_t a16 = 0;
	int32_t a17 = 0;
	int32_t a18 = 0;
	int32_t a19 = 0;
	int32_t a20 = 0;
	int32_t a21 = 0;
	int32_t a22 = 0;
	int32_t a23 = 0;
	int32_t a24 = 0;
	int32_t a25 = 0;
	int32_t a26 = 0;
	int32_t a27 = 0;
	int32_t a28 = 0;
	int32_t a29 = 0;
	int32_t a30 = 0;
	int32_t a31 = 0;

	std::string_view tl_name() const { return "tasks.fullFilledCron"; }
	uint32_t tl_tag() const { return 0xd4177d7e; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const FullFilledCron& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

