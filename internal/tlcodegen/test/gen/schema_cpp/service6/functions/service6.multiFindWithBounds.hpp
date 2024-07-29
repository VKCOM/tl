#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../types/service6.findWithBoundsResult.hpp"
#include "../../__common/types/Either.hpp"


namespace tl2 { namespace service6 { 
struct MultiFindWithBounds {
	std::vector<int32_t> clusters;

	std::string_view tl_name() const { return "service6.multiFindWithBounds"; }
	uint32_t tl_tag() const { return 0x84b168cf; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> & result);
	bool write_result(::basictl::tl_ostream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> & result);
};

}} // namespace tl2::service6

