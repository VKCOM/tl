#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { namespace tasks { 
struct QueueStats {
	int32_t waiting_size = 0;
	int32_t scheduled_size = 0;
	int32_t in_progress_size = 0;

	std::string_view tl_name() const { return "tasks.queueStats"; }
	uint32_t tl_tag() const { return 0x1d942543; }

	bool write_json(std::ostream& s, uint32_t nat_fields_mask)const;

	bool read(::basictl::tl_istream & s, uint32_t nat_fields_mask);
	bool write(::basictl::tl_ostream & s, uint32_t nat_fields_mask)const;

	bool read_boxed(::basictl::tl_istream & s, uint32_t nat_fields_mask);
	bool write_boxed(::basictl::tl_ostream & s, uint32_t nat_fields_mask)const;
};

}} // namespace tl2::tasks

