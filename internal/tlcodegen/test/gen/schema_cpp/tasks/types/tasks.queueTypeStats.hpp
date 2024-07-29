#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace tasks { 
struct QueueTypeStats {
	uint32_t fields_mask = 0;
	int64_t waiting_size = 0;
	int64_t scheduled_size = 0;
	int64_t in_progress_size = 0;
	int32_t num_queues = 0;

	std::string_view tl_name() const { return "tasks.queueTypeStats"; }
	uint32_t tl_tag() const { return 0xe1b785f2; }

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::tasks

