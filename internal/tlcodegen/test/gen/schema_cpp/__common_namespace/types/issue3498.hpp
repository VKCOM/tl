#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../service6/types/service6.findResultRow.hpp"
#include "../../service6/types/service6.error.hpp"
#include "Either.hpp"


namespace tl2 { 
struct Issue3498 {
	std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> x;

	std::string_view tl_name() const { return "issue3498"; }
	uint32_t tl_tag() const { return 0xf54b7b0a; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

} // namespace tl2

