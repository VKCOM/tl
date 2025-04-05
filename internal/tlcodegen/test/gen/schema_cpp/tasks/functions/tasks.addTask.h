#pragma once

#include "../../basictl/io_streams.h"
#include "../types/tasks.task.h"


namespace tl2 { namespace tasks { 
struct AddTask {
	std::string type_name;
	std::vector<int32_t> queue_id;
	::tl2::tasks::Task task{};

	std::string_view tl_name() const { return "tasks.addTask"; }
	uint32_t tl_tag() const { return 0x2ca073d5; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, int64_t & result);
	bool write_result(::basictl::tl_ostream & s, int64_t & result);

	friend std::ostream& operator<<(std::ostream& s, const AddTask& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

