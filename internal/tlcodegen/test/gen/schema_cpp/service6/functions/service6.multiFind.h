#pragma once

#include "../../basictl/io_streams.h"
#include "../types/service6.findResultRow.h"
#include "../types/service6.error.h"
#include "../../__common_namespace/types/Either.h"


namespace tl2 { namespace service6 { 
struct MultiFind {
	std::vector<int32_t> clusters;
	int32_t limit = 0;
	double eq_threshold = 0;

	std::string_view tl_name() const { return "service6.multiFind"; }
	uint32_t tl_tag() const { return 0xe62178d8; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result);

	friend std::ostream& operator<<(std::ostream& s, const MultiFind& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service6

