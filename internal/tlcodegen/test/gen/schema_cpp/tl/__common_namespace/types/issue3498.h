#pragma once

#include "../../../basics/basictl.h"
#include "../../service6/types/service6.findResultRow.h"
#include "../../service6/types/service6.error.h"
#include "Either.h"


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

	friend std::ostream& operator<<(std::ostream& s, const Issue3498& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

