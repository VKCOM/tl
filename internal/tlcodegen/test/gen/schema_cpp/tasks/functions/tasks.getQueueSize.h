#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/types/tasks.queueStats.h"


namespace tl2 { namespace tasks { 
struct GetQueueSize {
	std::string type_name;
	std::vector<int32_t> queue_id;
	uint32_t fields_mask = 0;
	int32_t local_dep = 0;

	std::string_view tl_name() const { return "tasks.getQueueSize"; }
	uint32_t tl_tag() const { return 0x6abbb057; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::tasks::QueueStats & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::tasks::QueueStats & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::tasks::QueueStats & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::tasks::QueueStats & result);

	friend std::ostream& operator<<(std::ostream& s, const GetQueueSize& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

