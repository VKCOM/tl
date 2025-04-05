#pragma once

#include "../../basictl/io_streams.h"
#include "antispam.patternNotFound.h"
#include "antispam.patternFound.h"


namespace tl2 { namespace antispam { 
struct PatternFull {
	std::variant<::tl2::antispam::PatternFound, ::tl2::antispam::PatternNotFound> value;

	bool is_patternFound() const { return value.index() == 0; }
	bool is_patternNotFound() const { return value.index() == 1; }

	void set_patternNotFound() { value.emplace<1>(); }

	std::string_view tl_name() const;
	uint32_t tl_tag() const;

	bool write_json(std::ostream& s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::antispam

