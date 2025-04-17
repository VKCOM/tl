#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tl2 { namespace service3 { 
struct DeleteAllProducts {
	int32_t user_id = 0;
	int32_t type = 0;
	int32_t start_date = 0;
	int32_t end_date = 0;

	std::string_view tl_name() const { return "service3.deleteAllProducts"; }
	uint32_t tl_tag() const { return 0x4494acc2; }

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

	void read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result);

	friend std::ostream& operator<<(std::ostream& s, const DeleteAllProducts& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

