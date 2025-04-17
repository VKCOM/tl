#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tl2 { namespace tasks { 
struct QueueTypeSettings {
	uint32_t fields_mask = 0;
	bool is_enabled = false;
	bool is_persistent = false;
	int32_t priority = 0;
	int32_t default_retry_time = 0;
	int32_t default_retry_num = 0;
	std::string move_to_queue_type_on_error;
	bool is_blocking = false;
	int32_t timelimit = 0;
	int32_t max_queue_size = 0;

	std::string_view tl_name() const { return "tasks.queueTypeSettings"; }
	uint32_t tl_tag() const { return 0x561fbc09; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const QueueTypeSettings& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

