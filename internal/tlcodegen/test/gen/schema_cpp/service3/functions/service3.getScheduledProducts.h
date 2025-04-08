#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service3.product.h"


namespace tl2 { namespace service3 { 
struct GetScheduledProducts {
	int32_t user_id = 0;
	std::vector<int32_t> types;

	std::string_view tl_name() const { return "service3.getScheduledProducts"; }
	uint32_t tl_tag() const { return 0xf53ad7bd; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>> & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>> & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>> & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>> & result);

	friend std::ostream& operator<<(std::ostream& s, const GetScheduledProducts& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service3

