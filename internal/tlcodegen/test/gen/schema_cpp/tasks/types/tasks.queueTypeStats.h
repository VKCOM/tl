#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tl2 { namespace tasks { 
struct QueueTypeStats {
	uint32_t fields_mask = 0;
	int64_t waiting_size = 0;
	int64_t scheduled_size = 0;
	int64_t in_progress_size = 0;
	int32_t num_queues = 0;

	std::string_view tl_name() const { return "tasks.queueTypeStats"; }
	uint32_t tl_tag() const { return 0xe1b785f2; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const QueueTypeStats& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

