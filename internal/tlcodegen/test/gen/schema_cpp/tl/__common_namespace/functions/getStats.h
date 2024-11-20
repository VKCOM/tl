#pragma once

#include "../../../basics/basictl.h"
#include "../../tasks/types/tasks.queueTypeStats.h"


namespace tl2 { 
struct GetStats {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0xbaa6da35; }

	::tl2::tasks::QueueTypeStats x{};

	std::string_view tl_name() const { return "getStats"; }
	uint32_t tl_tag() const { return 0xbaa6da35; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats & result);
	bool write_result(::basictl::tl_ostream & s, ::tl2::tasks::QueueTypeStats & result);

	friend std::ostream& operator<<(std::ostream& s, const GetStats& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

