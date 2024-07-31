#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "tasks.queueTypeSettings.hpp"
#include "tasks.queueTypeStats.hpp"


namespace tl2 { namespace tasks { 
struct QueueTypeInfo {
	std::string type_name;
	::tl2::tasks::QueueTypeSettings settings{};
	::tl2::tasks::QueueTypeStats stats{};

	std::string_view tl_name() const { return "tasks.queueTypeInfo"; }
	uint32_t tl_tag() const { return 0x38d38d3e; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const QueueTypeInfo& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

