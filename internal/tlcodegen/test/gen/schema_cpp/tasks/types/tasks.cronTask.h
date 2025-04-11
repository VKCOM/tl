#pragma once

#include "../../basictl/io_streams.h"
#include "tasks.cronTime.h"
#include "tasks.task.h"


namespace tl2 { namespace tasks { 
struct CronTask {
	std::string type_name;
	std::vector<int32_t> queue_id;
	::tl2::tasks::Task task{};
	::tl2::tasks::CronTime time{};

	std::string_view tl_name() const { return "tasks.cronTask"; }
	uint32_t tl_tag() const { return 0xc90cf28a; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const CronTask& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

