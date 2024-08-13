#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service6.findResultRow.hpp"
#include "../types/service6.error.hpp"
#include "../../__common_namespace/types/Either.hpp"


namespace tl2 { namespace service6 { 
struct MultiFind {
	// tl magic for function
	static const uint32_t MAGIC = 0xe62178d8;

	std::vector<int32_t> clusters;
	int32_t limit = 0;
	double eq_threshold = 0;

	std::string_view tl_name() const { return "service6.multiFind"; }
	uint32_t tl_tag() const { return 0xe62178d8; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result);

	friend std::ostream& operator<<(std::ostream& s, const MultiFind& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::service6

