#pragma once

#include "../../basictl/io_streams.h"
#include "service4.object.h"


namespace tl2 { namespace service4 { 
struct ModifiedNewsEntry {
	::tl2::service4::Object object{};
	int32_t creation_date = 0;
	uint32_t fields_mask = 0;
	int32_t restoration_date = 0;
	int32_t deletion_date = 0;
	bool hidden_by_privacy = false;

	std::string_view tl_name() const { return "service4.modifiedNewsEntry"; }
	uint32_t tl_tag() const { return 0xda19832a; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const ModifiedNewsEntry& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service4

